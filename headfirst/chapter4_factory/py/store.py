
from abc import ABC,abstractmethod
from pizzatypes import Margherita, VeggiePizza, VeganPizza
from pizza import Pizza
from factory import IndiaIngrediantFactory, NYIngredientFactory

class PizzaStore(ABC):
    @abstractmethod
    def create(self,type) -> Pizza:
        pass

    def order(self, type):
        pizza = self.create(type)
        pizza.prepare()
        pizza.bake()
        pizza.pack()

class IndiaPizzaStore(PizzaStore):

    def __init__(self) -> None:
        super().__init__()
        self.ingredientFactory = IndiaIngrediantFactory()

    def create(self, type) -> Pizza:
        if type=='veggie':
            return VeggiePizza(self.ingredientFactory)
        raise Exception("Only veg")
    
class NYPizzaStore(PizzaStore):
    def __init__(self) -> None:
        super().__init__()
        self.ingredientFactory = NYIngredientFactory()

    def create(self,type) -> Pizza:
        if type=='veggie':
            return VeganPizza(self.ingredientFactory)
        elif type=='Margherita':
            return Margherita(self.ingredientFactory)
        raise Exception("invalid type")


if __name__=='__main__':
    nps = NYPizzaStore()
    nps.order("veggie")

    ips = IndiaPizzaStore()
    ips.order("veggie")
