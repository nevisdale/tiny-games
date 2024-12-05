score = 0
rungame = false
gameover = false

function restore()
	bg_reset()
	player_reset()
	pipes_reset()
	enemies_reset()
	score = 0
	gameover = false
	rungame = false
end

function _init()
	restore()
end

function _update60()
	bg_update()

	if gameover then
		if btnp(5) then
			restore()
		end
		return
	end

	if btnp(5) then
		rungame = true
		player_up()
	end

	local dir = 0
	if btn(0) then
		dir = -1
	elseif btn(1) then
		dir = 1
	end
	player_movex(dir)

	if not rungame then
		return
	end

	player_update()
	pipes_update()
	enemies_update()

	-- pass pipes and increase the score
	foreach(
		pipes, function(p)
			if p.posx + p.w < player.posx and not p.passed then
				score += 1
				p.passed = true
			end
		end
	)

	-- CHECK GAME OVER
	gameover = gameover or player.posy < 0 or player.posy > 128 - player.h
	foreach(
		pipes, function(p)
			if player.posx + player.w >= p.posx and player.posx <= p.posx + p.w then
				gameover = gameover
						or player.posy < p.top
						or player.posy + player.w > p.btm
			end
		end
	)
end

function _draw()
	cls()
	bg_draw()
	pipes_draw()
	enemies_draw()
	player_draw()
	print("score: " .. score, 1, 1, 8)

	if gameover then
		rectfill(18, 60, 118, 72, 2)
		print("press any key to restart", 20, 64, 7)
	end
end