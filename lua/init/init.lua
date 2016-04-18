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

function Clawn:new()
    self.__index = self
    return setmetatable({}, self)
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

clawn_count = 80
clawn_factory = {}

for i = 1, clawn_count do
    clawn_factory[i] = Clawn:new()
    c = clawn_factory[i]
    c:setMaxLiveCycle(30)
end


Shark = {}

function Shark:new()
    self.__index = self
    return setmetatable({}, self)
end


function Shark:setMaxLiveCycle(max_live_cycle)
    if max_live_cycle == nil or max_live_cycle <= 0 then
        self.max_live_cycle = 0
        return
    end
    math.randomseed( os.time() )
    rand_delta_persent = math.random(3, 20)
    delta = max_live_cycle * rand_delta_persent / 100
    if math.random(2) == 1 then
        self.max_live_cycle = max_live_cycle + delta
    else
        self.max_live_cycle = max_live_cycle + delta
    end
end

shark_count = 20
shark_factory = {}

for i = 1, shark_count do
    shark_factory[i] = Shark:new()
    s = shark_factory[i]
    s:setMaxLiveCycle(70)
end
