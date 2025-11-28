package Contracts

import "Spacetraders/src/General"
import "encoding/json"

func UpdateContracts() error {
	jsonStr := General.GetUrlJson("https://api.spacetraders.io/v2/my/contracts", "agent")
	var wrapper map[string]json.RawMessage
	err := json.Unmarshal([]byte(jsonStr), &wrapper)
	if err != nil { General.LogErr("UpdateContracts: " + err.Error()) }

	var c []Contract
	err = json.Unmarshal(wrapper["data"], &c)
	if err != nil { General.LogErr("UpdateContracts: " + err.Error()) }

	for index := range c {
		_, err = General.PG.Exec(`
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
		if err != nil { General.LogErr("UpdateContracts contracts: " + err.Error()) }

		for idx2 := range c[index].Terms.Deliver {
			_, err = General.PG.Exec(`
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
		if err != nil { General.LogErr("UpdateContracts contractMats: " + err.Error()) }
		}
	}

	return nil
}

func GetContract(id string) Contract {
	var result Contract

	// Fill structs for each contract
	data, _ := General.PG.Query(`SELECT * FROM contracts WHERE id = $1`, id)
	for data.Next() {
		data.Scan(
			&result.ID,
			&result.Faction,
			&result.Type,
			&result.Terms.Payment.OnAccepted,
			&result.Terms.Payment.OnFulfilled,
			&result.Accepted,
			&result.Fulfilled,
			&result.Terms.Deadline,
			&result.Expiration,
			&result.DeadlineToAccept,
			&result.LastUpdated,
		)
	}

	mats, _ := General.PG.Query(`SELECT material,destination,units_required,units_fulfilled FROM contract_materials WHERE id = $1`, id)
	for mats.Next() {
		var deliverData ContractDeliveries
		mats.Scan(
			&deliverData.Material,
			&deliverData.Destination,
			&deliverData.UnitsRequired,
			&deliverData.UnitsFulfilled,
		)
		result.Terms.Deliver = append(result.Terms.Deliver, deliverData) //> Or you can append to it if you don't know how many elements there will be.
	}

	return result
}

func NegotiateNewContract(ship string) {
	General.PostUrlJson("https://api.spacetraders.io/v2/my/ships/"+ship+"/negotiate/contract", "agent")
}

func AcceptContract(contract string) {
	General.PostUrlJson("https://api.spacetraders.io/v2/my/contracts/"+contract+"/accept", "agent")
}
