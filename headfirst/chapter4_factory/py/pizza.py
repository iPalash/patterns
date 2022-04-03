from abc import ABC, abstractmethod
class Pizza(ABC):

    @abstractmethod
    def prepare(self):
        pass

    def bake(self):
        print("Baking "+ self.Name()) 

    def pack(self):
        print("Packing " +self.Name())

    def Name(self):
        return "{} with {} on {} with added {}".format(self.name,self.sauce,self.base, ','.join(map(str,self.toppings))) 
