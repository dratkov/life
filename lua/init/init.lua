Area = require("area")

area = Area:new(50, 40)

Cell = require('cell')

Clawn = require('clawn')

clawn_count = 80
clawn_factory = {}

for i = 1, clawn_count do
    clawn_factory[i] = Clawn:new()
    c = clawn_factory[i]
end


Shark = require('shark')

shark_count = 20
shark_factory = {}

for i = 1, shark_count do
    shark_factory[i] = Shark:new()
    s = shark_factory[i]
end
