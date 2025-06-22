from dataclasses import dataclass
from uuid import UUID

@dataclass
class Range:
    Min: int
    Max: int
    Normal: int
    Step: int = None

@dataclass
class Resolution:
    X: int
    Y: int = None

    def __post_init__(self):
        if self.Y == None:
            self.Y = self.X

    def __repr__(self):
        if self.X == self.Y:
            return 'Resolution({})'.format(self.X)

        return 'Resolution(X={}, Y={})'.format(self.X, self.Y)
