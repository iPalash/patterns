from abc import abstractmethod,ABC
from ingredients import Sauce, Base, SoySauce, ThickCrust, ThinCrust, TomatoSauce, Topping, Onion, Paneer, Tofu, Olive

            

class PizzaIngredientFactory(ABC):
    @abstractmethod
    def createSauce() -> Sauce:
        pass

    @abstractmethod
    def createBase() -> Base:
        pass

    @abstractmethod
    def createToppings() -> list[Topping]:
        pass

class IndiaIngrediantFactory(PizzaIngredientFactory):
    def createSauce(self) -> Sauce:
        return TomatoSauce()
    def createBase(self) -> Base:
        return ThickCrust()
    def createToppings(self) -> list[Topping]:
        return [Onion(), Paneer()]

class NYIngredientFactory(PizzaIngredientFactory):
    def createSauce(self) -> Sauce:
        return SoySauce()
    def createBase(self) -> Base:
        return ThinCrust()
    def createToppings(self) -> list[Topping]:
        return [Tofu(),Olive()]