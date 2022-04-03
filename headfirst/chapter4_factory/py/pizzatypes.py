from abc import ABC,abstractmethod
from pizza import Pizza
from factory import PizzaIngredientFactory


class VeggiePizza(Pizza):
    def __init__(self, ingredientFactory : PizzaIngredientFactory) -> None:
        super().__init__()
        self.ingredientFactory = ingredientFactory
        # self.base = "ThickCrust"
        # self.sauce = "TomatoSauce"
        # self.toppings = ["onion","paneer"]
        self.name = "VeggiePizza"
    
    def prepare(self):
        self.base = self.ingredientFactory.createBase()
        self.sauce = self.ingredientFactory.createSauce()
        self.toppings = self.ingredientFactory.createToppings()

class VeganPizza(Pizza):
    def __init__(self, ingredientFactory : PizzaIngredientFactory) -> None:
        super().__init__()
        self.ingredientFactory = ingredientFactory
        # self.base = "ThinCrust"
        # self.sauce = "SoySauce"
        # self.toppings = ["tofu","olive"]
        self.name = "VeganPizza"
    
    def prepare(self):
        self.base = self.ingredientFactory.createBase()
        self.sauce = self.ingredientFactory.createSauce()
        self.toppings = self.ingredientFactory.createToppings()

class Margherita(Pizza):
    def __init__(self, ingredientFactory: PizzaIngredientFactory) -> None:
        super().__init__()
        self.ingredientFactory = ingredientFactory
        # self.base = "ThinkCrust"
        # self.sauce = "TomatoSauce"
        # self.toppings = ["mozarella","cheddar"]
        self.name = "Margherita"
    def prepare(self):
        self.base = self.ingredientFactory.createBase()
        self.sauce = self.ingredientFactory.createSauce()
