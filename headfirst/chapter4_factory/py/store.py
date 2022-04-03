from abc import ABC

from factory import SimplePizzaFactory
from abc import ABC,abstractmethod
from pizza import Margherita, Pizza, VeggiePizza, VeganPizza


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
    def create(self, type) -> Pizza:
        if type=='veggie':
            return VeggiePizza()
        raise Exception("Only veg")
    
class NYPizzaStore(PizzaStore):
    def create(self,type) -> Pizza:
        if type=='veggie':
            return VeganPizza()
        elif type=='Margherita':
            return Margherita()
        raise Exception("invalid type")


if __name__=='__main__':
    nps = NYPizzaStore()
    nps.order("veggie")

    ips = IndiaPizzaStore()
    ips.order("veggie")
