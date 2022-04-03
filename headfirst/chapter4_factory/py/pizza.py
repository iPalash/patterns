from abc import ABC,abstractmethod


class Pizza(ABC):
    def bake(self):
        print("Baking "+ self.Name()) 

    def pack(self):
        print("Packing " +self.Name())

    def Name(self):
        return self.name


class VeggiePizza(Pizza):
    def __init__(self) -> None:
        super().__init__()
        self.name = "VeggiePizza"

class VeganPizza(Pizza):
    def __init__(self) -> None:
        super().__init__()
        self.name = "VeganPizza"

class Margherita(Pizza):
    def __init__(self) -> None:
        super().__init__()
        self.name = "Margherita"
