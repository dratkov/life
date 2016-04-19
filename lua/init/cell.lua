local Cell = {}

function Cell:new()
    self.__index = self
    return setmetatable({}, self)
end

return Cell
