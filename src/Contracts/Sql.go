package Contracts

import (
	"Spacetraders/src/General"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func GetAllContracts() []Contract {
	db, err := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	if err != nil {
		General.LogErr(fmt.Sprintf("DB open failed: %v", err))
	}
	defer db.Close()

	// Get all contract IDs
	ids, _ := db.Query(`SELECT id FROM contracts`)
	defer ids.Close()

	var contracts []string
	for ids.Next() {
		var s string
		ids.Scan(&s)
		contracts = append(contracts, s)
	}

	// Fill structs for each contract
	CStruct := make([]Contract, len(contracts)) //> You must instantiate the substruct index before you can .Scan() to it apparently
	for idx, value := range contracts {
		data, _ := db.Query(`SELECT * FROM contracts WHERE id = $1`, value)
		for data.Next() {
			row := &CStruct[idx]
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

		mats, _ := db.Query(`SELECT material,destination,units_required,units_fulfilled FROM contract_materials WHERE id = $1`, value)
		for mats.Next() {
			var deliverData ContractDeliveries
			mats.Scan(
				&deliverData.Material,
				&deliverData.Destination,
				&deliverData.UnitsRequired,
				&deliverData.UnitsFulfilled,
			)
			CStruct[idx].Terms.Deliver = append(CStruct[idx].Terms.Deliver, deliverData) //> Or you can append to it if you don't know how many elements there will be.
		}
	}

	return CStruct
}

func UpdateContracts() error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/my/contracts", "agent")
	var wrapper map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &wrapper)
	if err != nil {
		log.Fatal(err)
	}

	var c []Contract
	err = json.Unmarshal(wrapper["data"], &c)
	if err != nil {
		General.LogErr(fmt.Sprintf("%v", err))
	}

	db, err := sql.Open("postgres", "user=skyehunter dbname=spacetraders sslmode=disable")
	if err != nil {
		General.LogErr(fmt.Sprintf("DB open failed: %v", err))
		return err
	}
	defer db.Close()

	for index := range c {
		_, err = db.Exec(`
			INSERT INTO contracts (id,faction,type,deadline,pay_on_accept,pay_on_complete,accepted,fulfilled,expiration,deadline_to_accept,last_updated) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW())
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
			c[index].ID,
			c[index].Faction,
			c[index].Type,
			c[index].Terms.Deadline,
			c[index].Terms.Payment.OnAccepted,
			c[index].Terms.Payment.OnFulfilled,
			c[index].Accepted,
			c[index].Fulfilled,
			c[index].Expiration,
			c[index].DeadlineToAccept,
		)
		if err != nil {
			General.LogErr(fmt.Sprintf("%v", err))
		}

		for idx2 := range c[index].Terms.Deliver {
			_, err = db.Exec(`
				INSERT INTO contract_materials (id,material,destination,units_required,units_fulfilled) 
				VALUES ($1,$2,$3,$4,$5)
				ON CONFLICT (id,material,destination) DO UPDATE SET
					id              = EXCLUDED.id, 
					material        = EXCLUDED.material,
					destination     = EXCLUDED.destination,
					units_required  = EXCLUDED.units_required,
					units_fulfilled = EXCLUDED.units_fulfilled
				`,
				c[index].ID,
				c[index].Terms.Deliver[idx2].Material,
				c[index].Terms.Deliver[idx2].Destination,
				c[index].Terms.Deliver[idx2].UnitsRequired,
				c[index].Terms.Deliver[idx2].UnitsFulfilled,
			)
			if err != nil {
				General.LogErr(fmt.Sprintf("%v", err))
			}
		}
	}

	return nil
}

func NegotiateNewContract(ship string) {
	General.PostUrlJson("https://api.spacetraders.io/v2/my/ships/"+ship+"/negotiate/contract", "agent")
}

func AcceptContract(contract string) {
	General.PostUrlJson("https://api.spacetraders.io/v2/my/contracts/"+contract+"/accept", "agent")
}
