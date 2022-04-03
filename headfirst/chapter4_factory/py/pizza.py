from abc import ABC,abstractmethod


class Pizza(ABC):

    def prepare(self):
        print("using base: ", self.base)
        print("applying sauce: ",self.sauce)
        for topping in self.toppings:
            print("adding topping:",topping)

    def bake(self):
        print("Baking "+ self.Name()) 

    def pack(self):
        print("Packing " +self.Name())

    def Name(self):
        return "{} with {} on {} with added {}".format(self.name,self.sauce,self.base, ','.join(self.toppings)) 


class VeggiePizza(Pizza):
    def __init__(self) -> None:
        super().__init__()
        self.base = "ThickCrust"
        self.sauce = "TomatoSauce"
        self.toppings = ["onion","paneer"]
        self.name = "VeggiePizza"

class VeganPizza(Pizza):
    def __init__(self) -> None:
        super().__init__()
        self.base = "ThinCrust"
        self.sauce = "SoySauce"
        self.toppings = ["tofu","olive"]
        self.name = "VeganPizza"

class Margherita(Pizza):
    def __init__(self) -> None:
        super().__init__()
        self.base = "ThinkCrust"
        self.sauce = "TomatoSauce"
        self.toppings = ["mozarella","cheddar"]
        self.name = "Margherita"
