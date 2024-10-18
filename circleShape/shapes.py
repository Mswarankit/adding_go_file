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


class Triangle(Shape):
    def __init__(self, base, height):
        super().__init__(name="Triangle")
        self.base = base
        self.height = height
    
    def calculate_area(self):
        return (1/2)* self.base * self.height
    
    def calculate_volume(self):
        return 0

class Rectangle(Shape):
    def __init__(self, length, width):
        super().__init__(name="Rectangle")
        self.length = length
        self.width = width

    def calculate_area(self):
        return self.length * self.width
    
    def calculate_volume(self):
        return 0

class Square(Shape):
    def __init__(self, side):
        super().__init__(name="Square")
        self.side = side
    
    def calculate_area(self):
        return math.pow(self.side, 2)
    
    def calculate_volume(self):
        return 0
    

class Circle(Shape):
    def __init__(self, radius):
        super().__init__(name="Circle")
        self.radius =  radius

    def calculate_area(self):
        return math.pi * self.radius ** 2

    def calculate_volume(self):
        return 0

class Cube(Square):
    def __init__(self, side):
        super().__init__(side=side)
        self.name = "Cube"
    
    def calculate_area(self):
        return 6 * super().calculate_area()
    
    def calculate_volume(self):
        return math.pow(self.side, 3)

class Cuboid(Rectangle):
    def __init__(self, length, width, height):
        super().__init__(length, width)
        self.height = height
        self.name = "Cuboid"
    
    def calculate_area(self):
        return 2*(super().calculate_area() + (self.length * self.height) + (self.height * self.width))
    
    def calculate_volume(self):
        return super().calculate_area()*self.height

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
    
class Hemishpere(Circle):
    def __init__(self, radius):
        super().__init__(radius)
        self.name = "Hemishpere"
    
    def calculate_area(self):
        return 3 * super().calculate_area()
    
    def calculate_volume(self):
        return (2/3)*super().calculate_volume()
    
# 2d shapes
tri = Triangle(12, 7)
rect = Rectangle(15, 19)
squa = Square(16)
c = Circle(15)
# 3d shpaes
cyl = Cylinder(10, 12)
cone = Cone(12, 15)
sp = Sphere(5)
cu = Cube(6)
cuboid = Cuboid(12, 15, 11)
hsp = Hemishpere(4)


# shp = [tri, rect, squa, c, cyl, cone, sp, cu, cuboid, hsp]
# for i in shp:
#     i.printAll()


