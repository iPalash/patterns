package strategy1

// Character interface

type Character interface {
	fight()
}

type CharacterImpl struct {
	w WeaponBehaviour
}

func (c *CharacterImpl) SetWeapon(w WeaponBehaviour) {
	c.w = w
}

type King struct {
	CharacterImpl
}

func NewKing() Character {
	k := &King{CharacterImpl{}}
	k.SetWeapon(&SwordBehaviour{})
	return k
}

func (k *King) fight() {
	k.CharacterImpl.w.useWeapon()
}

type Magician struct {
	CharacterImpl
}

func NewMagician() Character {
	k := &Magician{CharacterImpl{}}
	k.SetWeapon(&WandBehaviour{})
	return k
}

func (k *Magician) fight() {
	k.CharacterImpl.w.useWeapon()
}
