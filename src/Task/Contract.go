package Task 

import "Spacetraders/src/Model"
import "encoding/json"

func UpdateContracts(c []Model.Contract) error {
	var err error
	if c == nil {
		jsonStr := GetUrlJson("https://api.spacetraders.io/v2/my/contracts", "agent")
		var wrapper map[string]json.RawMessage
		err := json.Unmarshal([]byte(jsonStr), &wrapper)
		if err != nil { LogErr("UpdateContracts: " + err.Error()) }

		err = json.Unmarshal(wrapper["data"], &c)
		if err != nil { LogErr("UpdateContracts: " + err.Error()) }
	}

	for index := range c {
		_, err = PG.Exec(`
			INSERT INTO contracts (id,faction,type,deadline,pay_on_accept,pay_on_complete,accepted,fulfilled,expiration,deadline_to_accept,last_updated) 
			VALUES (?,?,?,?,?,?,?,?,?,?,datetime('now'))
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
		if err != nil { LogErr("UpdateContracts contracts: " + err.Error()) }

		for idx2 := range c[index].Terms.Deliver {
			_, err = PG.Exec(`
				INSERT INTO contract_materials (id,material,destination,units_required,units_fulfilled) 
				VALUES (?,?,?,?,?)
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
		if err != nil { LogErr("UpdateContracts contractMats: " + err.Error()) }
		}
	}

	return nil
}

func GetContract(id string) Model.Contract {
	var result Model.Contract

	// Fill structs for each contract
	data, _ := PG.Query(`SELECT * FROM contracts WHERE id = ?`, id)
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

	mats, _ := PG.Query(`SELECT material,destination,units_required,units_fulfilled FROM contract_materials WHERE id = ?`, id)
	for mats.Next() {
		var deliverData Model.ContractDeliveries
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
	PostUrlJson("https://api.spacetraders.io/v2/my/ships/"+ship+"/negotiate/contract", "agent")
	UpdateContracts(nil)
}

func AcceptContract(contract string) {
	PostUrlJson("https://api.spacetraders.io/v2/my/contracts/"+contract+"/accept", "agent")
	UpdateContracts(nil)
}
