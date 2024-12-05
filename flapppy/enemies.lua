enemies = {}

local function new_enemy(posx)
	return {
		posx = posx,
		posy = -8,
		btn = 10 + rnd(128 - 10),
		top = 0,
		spr = 16,
		spdy = 0.5
	}
end

function enemies_reset()
	enemies = {}
	add(enemies, new_enemy(64))
	add(enemies, new_enemy(96))
end

function enemies_update()
	-- TODO: pass score as an argument
	if score < 2 then
		return
	end

	foreach(
		enemies, function(e)
			e.posy += e.spdy
			if e.posy >= e.btn then
				e.spdy = -e.spdy
				e.posy = e.btn
				return
			end
			if e.posy < e.top and e.spdy < 0 then
				e.spdy = -e.spdy
				e.posy = 0
				e.btn = 10 + rnd(128 - 10)
			end
		end
	)
end

function enemies_draw()
	foreach(
		enemies, function(e)
			line(e.posx + 4, 0, e.posx + 4, e.posy)
			spr(e.spr, e.posx, e.posy)
		end
	)
end