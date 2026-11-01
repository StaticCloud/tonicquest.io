package combat

type Inventory struct {
	Item     Item
	Quantity int
}

type Item struct {
	Name   string
	Effect func()
}

type ItemManager struct {
	Player *Player
	Potion *Item
}

func InitItemManager(player *Player) *ItemManager {
	itemManager := &ItemManager{
		Player: player,
		Potion: &Item{
			Name:   "Potion",
			Effect: player.Reset,
		},
	}

	return itemManager
}

func (m *ItemManager) InitPotion() *Item {
	return m.Potion
}

func (i *Inventory) Use() {

	if i.Quantity > 0 {
		i.Item.Effect()
		i.Quantity -= 1
	}
}
