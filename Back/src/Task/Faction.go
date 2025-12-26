package Task

import "Spacetraders/src/Model"

func UpdateFaction(f *Model.Faction) error {
	_, err := PG.Exec(`DELETE FROM factions WHERE symbol = ?`, f.Symbol)
	if err != nil { LogErr("UpdateFaction: Failed to delete: " + err.Error()); return err }

	_, err = PG.Exec(`
		INSERT INTO factions (symbol, name, description, hq, recruiting)
		VALUES (?,?,?,?,?)
		ON CONFLICT (symbol) DO UPDATE SET
			symbol      = EXCLUDED.symbol,
			name        = EXCLUDED.name,
			description = EXCLUDED.description,
			hq          = EXCLUDED.hq,
			recruiting  = EXCLUDED.recruiting
		`,
		f.Symbol,
		f.Name,
		f.Description,
		f.HQ,
		f.Recruiting,
	)
	if err != nil { LogErr("UpdateFaction: Failed to insert: " + err.Error()); return err }

	for idx := range f.Traits {
		_, err = PG.Exec(`INSERT INTO faction_traits (faction, symbol, name, description) VALUES (?,?,?,?)`, f.Symbol, f.Traits[idx].Symbol, f.Traits[idx].Name, f.Traits[idx].Description)
	}

	return nil
}
