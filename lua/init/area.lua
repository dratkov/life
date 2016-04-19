local Area = {}

function Area:new(width, height)
    o = {}
    self.__index = self
    self.width = width or 0
    self.height = height or 0
    return setmetatable(o, self)
end

return Area