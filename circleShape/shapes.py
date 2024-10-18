import math

class Shape:
    def __init__(self, name):
        self.name = name

    def calculate_area(self):
        raise NotImplemented("Will implement in shape class")
    
    def calculate_volume(self):
        raise NotImplemented("will implement in different class")
    
    def printAll(self):
        print(f"Shape: {self.name}\nArea: {self.calculate_area():.2f}\nVolume: {self.calculate_volume():.2f}\n")



class Circle(Shape):
    def __init__(self, radius):
        super().__init__(name="Circle")
        self.radius =  radius

    def calculate_area(self):
        return math.pi * self.radius ** 2

    def calculate_volume(self):
        return 0
    


class Cylinder(Circle):
    def __init__(self, radius,height):
        super().__init__(radius)
        self.height = height
        self.name = "Cylinder"

    def calculate_volume(self):
        return math.pi * (self.radius ** 2) * self.height
    
    def calculate_area(self):
        return 2 * super().calculate_area() + 2 * math.pi * self.radius * self.height
    

class Cone(Cylinder):
    def __init__(self, radius, height):
        super().__init__(radius, height)
        self.name = "Cone"
    
    def calculate_area(self):
        t = math.pow
        lateral = t((t(self.radius,2)+t(self.height, 2)), 0.5)
        return math.pi * t(self.radius, 2) + math.pi * self.radius * lateral

    def calculate_volume(self):
        return (1/3)*super().calculate_volume()
    

class Sphere(Circle):
    def __init__(self, radius):
        super().__init__(radius)
        self.name = "Sphere"
    
    def calculate_area(self):
        return 4 * super().calculate_area()
    
    def calculate_volume(self):
        return (4/3)*super().calculate_area() * self.radius
    

c = Circle(15)
cyl = Cylinder(10, 12)
cone = Cone(12, 15)
sp = Sphere(5)


# shp = [c, cyl, cone, sp]
# for i in shp:
#     i.printAll()


