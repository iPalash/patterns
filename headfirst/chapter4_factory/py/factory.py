from pizza import Pizza, VeggiePizza, Margherita

class SimplePizzaFactory:
    def createPizza(self, type) -> Pizza:
        if type=='veggie':
            return VeggiePizza()
        elif type=='margherita':
            return Margherita()
        raise Exception("can't order that type")
            