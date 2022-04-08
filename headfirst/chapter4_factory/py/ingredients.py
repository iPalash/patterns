from abc import ABC,abstractmethod

class Sauce:
    def Name(self):
        return self.name
    def __repr__(self) -> str:
        return self.Name()
class TomatoSauce(Sauce):
    def __init__(self) -> None:
        super().__init__()
        self.name = "TomatoSauce"

class SoySauce(Sauce):
    def __init__(self) -> None:
        super().__init__()
        self.name = "SoySauce"

    
class Base:
    def Name(self):
        return self.name
    def __repr__(self) -> str:
        return self.Name()
class ThickCrust(Base):
    def __init__(self) -> None:
        super().__init__()
        self.name = "ThickCrust"

class ThinCrust(Base):
    def __init__(self) -> None:
        super().__init__()
        self.name = "ThinCrust"

class Topping:
    def Name(self):
        return self.name
    def __repr__(self) -> str:
        return self.Name()
class Onion(Topping):
    def __init__(self) -> None:
        super().__init__()
        self.name = "Onion"

class Paneer(Topping):
    def __init__(self) -> None:
        super().__init__()
        self.name = "Paneer"

class Tofu(Topping):
    def __init__(self) -> None:
        super().__init__()
        self.name = "Tofu"

class Olive(Topping):
    def __init__(self) -> None:
        super().__init__()
        self.name = "Olive"
