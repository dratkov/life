local Cell = require('cell')

local Clawn = Cell:new()

function Clawn:new()
    self.__index = self
    self.max_live_cycle = 30
    return setmetatable({}, self)
end

return Cell