player = {
	posx = 0,
	posy = 0,
	grv = 0,
	spr = 0,
	w = 0,
	h = 0
}

local sparks = {}

function player_up()
	player.grv = -2
end

function player_movex(dir)
	player.posx += dir
end

function player_update()
	player.posy += player.grv
	player.grv += 0.15

	add(
		sparks, {
			posx = player.posx + 1,
			posy = player.posy + rnd(player.h),
			clr = rnd({ 7, 8, 9, 14, 10 }),
			r = 1
		}
	)

	foreach(
		sparks, function(s)
			s.posx -= 2
			s.r += 0.25
			if s.posx < 0 then
				del(sparks, s)
			end
		end
	)
end

function player_draw()
	foreach(
		sparks, function(s)
			circ(s.posx, s.posy, s.r, s.clr)
		end
	)
	spr(player.spr, player.posx, player.posy)
end

function player_reset()
	player.posx = 32
	player.posy = 64
	player.grv = 1
	player.spr = 1
	player.w = 8
	player.h = 8
	sparks = {}
end