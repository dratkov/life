Area = {}

function Area:new(o, width, height)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    self.width = width or 0
    self.height = height or 0

    return o
end

area = Area:new(nil, 50, 40)

Clawn = {}

function Clawn:new(o, count, max_live_cycle)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    self.count = count or 0
    self:setMaxLiveCycle(max_live_cycle)

    return o
end

function Clawn:setMaxLiveCycle(max_live_cycle)
    if max_live_cycle == nil or max_live_cycle <= 0 then
        self.max_live_cycle = 0
        return
    end
    math.randomseed( os.time() )
    rand_delta_persent = math.random(5, 15)
    delta = max_live_cycle * rand_delta_persent / 100
    if math.random(2) == 1 then
        self.max_live_cycle = max_live_cycle + delta
    else
        self.max_live_cycle = max_live_cycle + delta
    end
end

clawn = Clawn:new(nil, 70)

Shark = {}

function Shark:new(o, count)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    self.count = count or 0

    return o
end

shark = Shark:new(nil, 40)