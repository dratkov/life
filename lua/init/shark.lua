local Shark = require('cell')

local Shark = Cell:new()

function Shark:new()
    self.__index = self
    self.max_live_cycle = 70
    return setmetatable({}, self)
end

return Shark
