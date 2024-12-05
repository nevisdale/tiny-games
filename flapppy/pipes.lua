pipes = {}

local top_btm_dist = 48
local pipe_min_size = 16

local function fill_pipe(pipe, posx)
	pipe.posx = posx
	pipe.w = 16
	pipe.dst = top_btm_dist
	pipe.top = pipe_min_size + rnd(64 - pipe_min_size)
	pipe.btm = pipe.top + pipe.dst
	pipe.spd = -3
	pipe.passed = false

	return pipe
end

function pipes_update()
	foreach(
		pipes, function(p)
			p.posx += p.spd
			-- TODO: pass score as an argument
			if score < 100 and p.posx + p.w < 0 then
				fill_pipe(p, 128)
			end
		end
	)
end

function pipes_draw()
	foreach(
		pipes, function(p)
			rectfill(p.posx, 0, p.posx + p.w, p.top, 3)
			rectfill(p.posx + 2, 0, p.posx + p.w - 2, p.top - 1, 1)
			rectfill(p.posx, p.btm, p.posx + p.w, 127, 3)
			rectfill(p.posx + 2, p.btm + 1, p.posx + p.w - 2, 127, 1)
		end
	)
end

function pipes_reset()
	pipes = {}
	add(pipes, fill_pipe({}, 110))
	add(pipes, fill_pipe({}, 110 + 64))
end