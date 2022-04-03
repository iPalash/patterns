package strategy1

import "fmt"

type WeaponBehaviour interface {
	useWeapon()
}

type SwordBehaviour struct {
}

func (s *SwordBehaviour) useWeapon() {
	fmt.Println("Attacking with Sword")
}

type WandBehaviour struct {
}

func (s *WandBehaviour) useWeapon() {
	fmt.Println("Casting a spell")
}
