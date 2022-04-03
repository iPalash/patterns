from abc import ABC

from factory import SimplePizzaFactory


class PizzaStore:

    def __init__(self, factory : SimplePizzaFactory) -> None:
        self.factory = factory

    def order(self, type):
        pizza = self.factory.createPizza(type)

        pizza.bake()
        pizza.pack()


if __name__=='__main__':
    ps = PizzaStore(SimplePizzaFactory())
    ps.order("veggie")
    ps.order("margherita")
