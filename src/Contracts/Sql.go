package Contracts

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func GetAllContracts() Contracts {
	db, _ := sql.Open("sqlite3", "SpaceTraders.db")
	defer db.Close()

	// Get all contract IDs
	ids, _ := db.Query(`SELECT id FROM contract`)
	defer ids.Close()

	var contracts []string
	for ids.Next() {
		var s string
		ids.Scan(&s)
		contracts = append(contracts, s)
	}

	// Fill structs for each contract
	var CStruct Contracts
	CStruct.Data = make([]ContractData, len(contracts)) //> You must instantiate the substruct index before you can .Scan() to it apparently
	for idx, value := range contracts {
		data, _ := db.Query(`SELECT * FROM contract WHERE id = ?`, value)
		for data.Next() {
			row := &CStruct.Data[idx]
			data.Scan(
				&row.ID,
				&row.Faction,
				&row.Type,
				&row.Terms.Payment.OnAccepted,
				&row.Terms.Payment.OnFulfilled,
				&row.Accepted,
				&row.Fulfilled,
				&row.Terms.Deadline,
				&row.Expiration,
				&row.DeadlineToAccept,
				&row.LastUpdated,
			)
		}

		mats, _ := db.Query(`SELECT material,destination,units_required,units_fulfilled FROM contract_materials WHERE id = ?`, value)
		for mats.Next() {
			var deliverData ContractDeliveries
			mats.Scan(
				&deliverData.Material,
				&deliverData.Destination,
				&deliverData.UnitsRequired,
				&deliverData.UnitsFulfilled,
			)
			CStruct.Data[idx].Terms.Deliver = append(CStruct.Data[idx].Terms.Deliver, deliverData) //> Or you can append to it if you don't know how many elements there will be.
		}
	}

	return CStruct
}

func UpdateContracts() error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/my/contracts", "agent")
	var c Contracts
	err := json.Unmarshal([]byte(jsonStr), &c)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := sql.Open("sqlite3", "SpaceTraders.db")
	defer db.Close()

	for index := range c.Data {
		_, err = db.Exec(`
			INSERT INTO contract (id,faction,type,deadline,pay_on_accept,pay_on_complete,accepted,fulfilled,expiration,deadline_to_accept,last_updated) 
			VALUES (?,?,?,?,?,?,?,?,?,?,datetime('now', 'localtime'))
			ON CONFLICT (id) DO UPDATE SET
				id                 = EXCLUDED.id, 
				faction            = EXCLUDED.faction, 
				type               = EXCLUDED.type, 
				deadline           = EXCLUDED.deadline, 
				pay_on_accept      = EXCLUDED.pay_on_accept, 
				pay_on_complete    = EXCLUDED.pay_on_complete, 
				accepted           = EXCLUDED.accepted, 
				fulfilled          = EXCLUDED.fulfilled, 
				expiration         = EXCLUDED.expiration, 
				deadline_to_accept = EXCLUDED.deadline_to_accept,
				last_updated       = EXCLUDED.last_updated
			`,
			c.Data[index].ID,
			c.Data[index].Faction,
			c.Data[index].Type,
			c.Data[index].Terms.Deadline,
			c.Data[index].Terms.Payment.OnAccepted,
			c.Data[index].Terms.Payment.OnFulfilled,
			c.Data[index].Accepted,
			c.Data[index].Fulfilled,
			c.Data[index].Expiration,
			c.Data[index].DeadlineToAccept,
		)
		if err != nil {
			General.LogErr(fmt.Sprintf("%v", err))
		}

		for idx2 := range c.Data[index].Terms.Deliver {
			_, err = db.Exec(`
				INSERT INTO contract_materials (id,material,destination,units_required,units_fulfilled) 
				VALUES (?,?,?,?,?)
				ON CONFLICT (id,material,destination) DO UPDATE SET
					id              = EXCLUDED.id, 
					material        = EXCLUDED.material,
					destination     = EXCLUDED.destination,
					units_required  = EXCLUDED.units_required,
					units_fulfilled = EXCLUDED.units_fulfilled
				`,
				c.Data[index].ID,
				c.Data[index].Terms.Deliver[idx2].Material,
				c.Data[index].Terms.Deliver[idx2].Destination,
				c.Data[index].Terms.Deliver[idx2].UnitsRequired,
				c.Data[index].Terms.Deliver[idx2].UnitsFulfilled,
			)
			if err != nil {
				General.LogErr(fmt.Sprintf("%v", err))
			}
		}
	}

	return nil
}

func NegotiateNewContract(ship string) {
	json := General.PostUrlJson("https://api.spacetraders.io/v2/my/ships/"+ship+"/negotiate/contract", "agent")
	General.LogActivity(string(json))
}
