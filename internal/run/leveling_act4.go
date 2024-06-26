package run

import (
	"github.com/hectorgimenez/d2go/pkg/data"
	"github.com/hectorgimenez/d2go/pkg/data/area"
	"github.com/hectorgimenez/d2go/pkg/data/npc"
	"github.com/hectorgimenez/d2go/pkg/data/quest"
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/game"
)

func (a Leveling) act4() action.Action {
	running := false
	return action.NewChain(func(d game.Data) []action.Action {
		if running || d.PlayerUnit.Area != area.ThePandemoniumFortress {
			return nil
		}

		running = true

		if !d.Quests[quest.Act4TheFallenAngel].Completed() {
			return a.izual()
		}

		return Diablo{baseRun: a.baseRun, bm: a.bm}.BuildActions()
	})
}

func (a Leveling) izual() []action.Action {
	return []action.Action{
		a.builder.MoveToArea(area.OuterSteppes),
		a.builder.Buff(),
		a.builder.MoveToArea(area.PlainsOfDespair),
		a.builder.Buff(),
		a.builder.MoveTo(func(d game.Data) (data.Position, bool) {
			izual, found := d.NPCs.FindOne(npc.Izual)
			if !found {
				return data.Position{}, false
			}

			return izual.Positions[0], true
		}),
		a.char.KillIzual(),
		a.builder.ReturnTown(),
		a.builder.InteractNPC(npc.Tyrael2),
	}
}
