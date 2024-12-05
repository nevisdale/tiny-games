local max_count = 120

local bg = {}

local function fillstar(s)
	s.posx = rnd(0)
	s.posy = rnd(128)
	s.spd = (rnd() + 0.5) * 1
	s.clr = rnd({ 1, 6, 7, 13 })
	return s
end

function bg_update()
	if count(bg) < max_count then
		add(bg, fillstar({}))
	end

	foreach(
		bg, function(s)
			s.posx += s.spd
			if s.posx > 128 then
				fillstar(s)
			end
		end
	)
end

function bg_draw()
	foreach(
		bg, function(s)
			pset(s.posx, s.posy, s.clr)
		end
	)
end

function bg_reset()
	for i = 1, max_count do
		bg_update()
	end
end