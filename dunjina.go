package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	// zoom
	zoom1, zoom2, zoom4 bool
	// menus
	menu1x, menu1y       int32
	menu1moveon, menu1on bool
	// maps
	obstructionmap = make([]string, levela)
	// select
	selectblock                = -1
	selectblockh, selectblockv int
	// player
	player, playernext                         = 4012, 4012
	playerh, playerv, playernexth, playernextv int
	// mouse
	mouseblock = -1
	// level
	levelw   = 1000
	levelh   = 1000
	levela   = levelh * levelw
	levelmap = make([]string, levela)
	// core
	drawblock, drawblocknext, drawa, drawblockv, drawblockh    int
	monh32, monw32                                             int32
	monitorh, monitorw, monitornum, blocksw, blocksh, blocknum int
	grid16on, grid4on, debugon, lrg, sml                       bool
	framecount                                                 int
	mousepos                                                   rl.Vector2
	camera                                                     rl.Camera2D
)

func timers() { // MARK: timers

}
func getpositions() { // MARK:getpositions()
	// horizontal vertical
	drawblockh = drawblocknext / levelw
	drawblockv = drawblocknext - (drawblockh * levelw)
	selectblockh = selectblock / levelw
	selectblockv = selectblock - (selectblockh * levelw)
	playerh = player / levelw
	playerv = player - (playerh * levelw)
	playernexth = playernext / levelw
	playernextv = playernext - (playernexth * levelw)

	// mouse block position
	xchange := float32(0)
	ychange := float32(0)
	ycount := 0
	if zoom1 {
		for b := 0; b < blocksh; b++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 16+ychange {
				for a := 0; a < blocksw; a++ {
					if mousepos.X > 0+xchange && mousepos.X < 16+xchange {
						mouseblock = a + ycount + drawblocknext
					}
					xchange += 16
				}
			}
			ychange += 16
			ycount += levelw
		}
	} else if zoom2 {
		for b := 0; b < blocksh/2; b++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 32+ychange {
				for a := 0; a < blocksw/2; a++ {
					if mousepos.X > 0+xchange && mousepos.X < 32+xchange {
						mouseblock = a + ycount + drawblocknext
					}
					xchange += 32
				}
			}
			ychange += 32
			ycount += levelw
		}
	} else if zoom4 {
		for b := 0; b < blocksh/4; b++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 64+ychange {
				for a := 0; a < blocksw/4; a++ {
					if mousepos.X > 0+xchange && mousepos.X < 64+xchange {
						mouseblock = a + ycount + drawblocknext
					}
					xchange += 64
				}
			}
			ychange += 64
			ycount += levelw
		}
	}

}
func screenposition() { // MARK: screenposition()

	if zoom1 {
		if playerh-drawblockh < 33 {
			if drawblockh > 0 {
				drawblocknext -= levelw
			}
		} else if playerh-drawblockh > 33 {
			if drawblockh < levelh-(blocksh+1) {
				drawblocknext += levelw
			}
		}

		if playerv-drawblockv < 60 {
			if drawblockv > 0 {
				drawblocknext--

			}
		} else if playerv-drawblockv > 60 {
			if drawblockv < levelw-(blocksw+1) {
				drawblocknext++

			}
		}
	} else if zoom2 {
		if playerh-drawblockh < 16 {
			if drawblockh > 0 {
				drawblocknext -= levelw

			}
		} else if playerh-drawblockh > 16 {
			if drawblockh < levelh-(blocksh+1) {
				drawblocknext += levelw

			}
		}

		if playerv-drawblockv < 30 {
			if drawblockv > 0 {
				drawblocknext--

			}
		} else if playerv-drawblockv > 30 {
			if drawblockv < levelw-(blocksw+1) {
				drawblocknext++

			}
		}
	} else if zoom4 {
		if playerh-drawblockh < 8 {
			if drawblockh > 0 {
				drawblocknext -= levelw

			}
		} else if playerh-drawblockh > 8 {
			if drawblockh < levelh-(blocksh+1) {
				drawblocknext += levelw

			}
		}

		if playerv-drawblockv < 15 {
			if drawblockv > 0 {
				drawblocknext--

			}
		} else if playerv-drawblockv > 15 {
			if drawblockv < levelw-(blocksw+1) {
				drawblocknext++

			}
		}
	}

}
func updateall() { // MARK: updateall()

	getpositions()
	screenposition()
	moveplayer()
	menus()

	if grid16on {
		grid16()
	}
	if grid4on {
		grid4()
	}
	timers()
}
func moveplayer() { // MARK: moveplayer()
	if playernext != player {

		if playernexth > playerh {
			if obstructionmap[player+levelw] == " " {
				player += levelw
			}
		} else if playernexth < playerh {
			if obstructionmap[player-levelw] == " " {
				player -= levelw
			}
		}
		if playernextv > playerv {
			if obstructionmap[player+1] == " " {
				player++
			}
		} else if playernextv < playerv {
			if obstructionmap[player-1] == " " {
				player--
			}
		}

	}
}
func menus() {

}
func createlevel() { // MARK: createlevel()

	for a := 0; a < levela; a++ {
		levelmap[a] = "."
		obstructionmap[a] = "."
	}

	levelmap[1] = "#"
	levelmap[3] = "#"
	levelmap[7] = "#"

	roomblock := 2010

	rooml := rInt(15, 25)
	roomw := rInt(15, 25)
	rooma := rooml * roomw
	count := 0

	for b := 0; b < 5; b++ {

		for a := 0; a < rooma; a++ {
			levelmap[roomblock] = "^"
			obstructionmap[roomblock] = " "
			roomblock++
			count++
			if count == rooml {
				count = 0
				roomblock += levelw - rooml
			}
		}
		roomblock += rooml
		roomblock -= rInt(4, 7) * levelw

		count = 0
		passagel := rInt(5, 10)
		passagew := rInt(2, 5)
		passagea := passagel * passagew

		for a := 0; a < passagea; a++ {
			levelmap[roomblock] = "^"
			obstructionmap[roomblock] = " "
			roomblock++
			count++
			if count == passagel {
				count = 0
				roomblock += levelw - passagel
			}
		}
		roomblock -= rInt(4, 7) * levelw
		rooml = rInt(15, 25)
		roomw = rInt(15, 25)
		rooma = rooml * roomw
		count = 0
	}

}
func startgame() { // MARK: startgame()
	createlevel()
}
func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window
	startsettings()
	raylib()
}
func raylib() { // MARK: raylib()
	rl.InitWindow(monw32, monh32, "dunjina")
	setscreen()
	startgame()
	rl.CloseWindow()
	rl.InitWindow(monw32, monh32, "dunjina")
	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	//	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() { // MARK: WindowShouldClose
		mousepos = rl.GetMousePosition()
		framecount++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// rl.DrawTexture(backimg, 0, 0, rl.Red) // MARK: draw backimg
		rl.BeginMode2D(camera)
		// MARK: draw map layer 1
		drawblock = drawblocknext
		linecount := 0
		drawx := int32(0)
		drawy := int32(0)
		for a := 0; a < blocknum; a++ {
			checklevel := levelmap[drawblock]
			switch checklevel {
			case ".":
			//	rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Brown, 0.1))
			case "#":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Magenta)
			case "^":
				rl.DrawRectangleLines(drawx, drawy, 15, 15, rl.Fade(rl.Orange, 0.2))
			}
			// draw mouseblock
			if mouseblock == drawblock {
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Orange)
			}
			// draw player
			if player == drawblock {
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Red)
			}
			// draw playernext
			if playernext == drawblock {
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Blue)
			}

			linecount++
			drawblock++
			drawx += 16
			if linecount == blocksw {
				linecount = 0
				drawx = 0
				drawy += 16
				drawblock += levelw - blocksw
			}
		}
		// MARK: draw map layer 2
		rl.EndMode2D() // MARK: draw no camera
		// draw vector no camera
		// draw menu1
		if menu1on {
			drawx = menu1x
			drawy = menu1y
			if menu1moveon {
				rl.DrawRectangle(drawx-157, drawy+3, 160, 320, rl.Fade(rl.White, 0.2))
				rl.DrawRectangle(drawx-160, drawy, 160, 320, rl.Black)
				rl.DrawRectangleLines(drawx-160, drawy, 160, 320, rl.Fade(rl.White, 0.2))
				rl.DrawRectangle(drawx, drawy, 16, 16, rl.Fade(rl.White, 0.2))

				rl.DrawText("click the white box right", drawx-150, drawy+10, 10, rl.White)
				rl.DrawText("to move this text box", drawx-150, drawy+20, 10, rl.White)
				rl.DrawText("left click new position", drawx-150, drawy+30, 10, rl.White)
				rl.DrawText("+/- keypad change zoom", drawx-150, drawy+40, 10, rl.White)
				rl.DrawText("del keypad no debug", drawx-150, drawy+50, 10, rl.White)
				rl.DrawText("F1/F2 draw grid", drawx-150, drawy+60, 10, rl.White)
			} else {
				rl.DrawRectangle(drawx-157, drawy+3, 160, 320, rl.Fade(rl.White, 0.7))
				rl.DrawRectangle(drawx-160, drawy, 160, 320, rl.Black)
				rl.DrawRectangleLines(drawx-160, drawy, 160, 320, rl.Fade(rl.White, 0.7))
				rl.DrawRectangle(drawx, drawy, 16, 16, rl.Fade(rl.White, 0.7))

				rl.DrawText("click the white box right", drawx-150, drawy+10, 10, rl.White)
				rl.DrawText("to move this text box", drawx-150, drawy+20, 10, rl.White)
				rl.DrawText("left click new position", drawx-150, drawy+30, 10, rl.White)
				rl.DrawText("+/- keypad change zoom", drawx-150, drawy+40, 10, rl.White)
				rl.DrawText("del keypad no debug", drawx-150, drawy+50, 10, rl.White)
				rl.DrawText("F1/F2 draw grid", drawx-150, drawy+60, 10, rl.White)
			}
		}
		// draw grid no camera
		drawblock = drawblocknext
		linecount = 0
		drawx = int32(0)
		drawy = int32(0)
		for a := 0; a < blocknum; a++ {
			linecount++
			drawblock++
			drawx += 16
			if linecount == blocksw {
				linecount = 0
				drawx = 0
				drawy += 16
				drawblock += levelw - blocksw
			}
		}

		if debugon {
			debug()
		}

		rl.EndDrawing()
		input()
		updateall()
	}
	rl.CloseWindow()
}
func setscreen() { // MARK: setscreen()
	monitornum = rl.GetMonitorCount()
	monitorh = rl.GetScreenHeight()
	monitorw = rl.GetScreenWidth()
	monh32 = int32(monitorh)
	monw32 = int32(monitorw)
	rl.SetWindowSize(monitorw, monitorh)
	setsizes()
}
func setsizes() { // MARK: setsizes()
	if monitorw >= 1600 {
		lrg = true
		sml = false
	} else if monitorw < 1600 && monitorw >= 1280 {
		lrg = false
		sml = true
	}
	blocksw = (monitorw / 16) + 1
	blocksh = (monitorh / 16) + 1
	blocknum = blocksh * blocksw
}
func startsettings() { // MARK: start
	camera.Zoom = 1.0
	zoom1 = true
	camera.Target.X = 0.0
	camera.Target.Y = 0.0
	debugon = true
	menu1x = 1300
	menu1y = 60
	menu1on = true
	//grid16on = true
	//selectedmenuon = true
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw32-300, 0, 500, monw32, rl.Fade(rl.Black, 0.7))
	rl.DrawFPS(monw32-290, monh32-100)

	monitorwTEXT := strconv.Itoa(monitorw)
	monitorhTEXT := strconv.Itoa(monitorh)
	blockswTEXT := strconv.Itoa(blocksw)
	blockshTEXT := strconv.Itoa(blocksh)
	mouseposXTEXT := fmt.Sprintf("%.0f", mousepos.X)
	mouseposYTEXT := fmt.Sprintf("%.0f", mousepos.Y)
	drawblockvTEXT := strconv.Itoa(drawblockv)
	drawblockhTEXT := strconv.Itoa(drawblockh)
	blocknumTEXT := strconv.Itoa(blocknum)
	mouseblockTEXT := strconv.Itoa(mouseblock)
	selectblockhTEXT := strconv.Itoa(selectblockh)
	selectblockvTEXT := strconv.Itoa(selectblockv)

	rl.DrawText(monitorwTEXT, monw32-290, 10, 10, rl.White)
	rl.DrawText("monitorw", monw32-200, 10, 10, rl.White)
	rl.DrawText(monitorhTEXT, monw32-290, 20, 10, rl.White)
	rl.DrawText("monitorh", monw32-200, 20, 10, rl.White)
	rl.DrawText(blockswTEXT, monw32-290, 30, 10, rl.White)
	rl.DrawText("blocksw", monw32-200, 30, 10, rl.White)
	rl.DrawText(blockshTEXT, monw32-290, 40, 10, rl.White)
	rl.DrawText("blocksh", monw32-200, 40, 10, rl.White)
	rl.DrawText(mouseposXTEXT, monw32-290, 50, 10, rl.White)
	rl.DrawText("mouseposX", monw32-200, 50, 10, rl.White)
	rl.DrawText(mouseposYTEXT, monw32-290, 60, 10, rl.White)
	rl.DrawText("mouseposY", monw32-200, 60, 10, rl.White)
	rl.DrawText(drawblockvTEXT, monw32-290, 70, 10, rl.White)
	rl.DrawText("drawblockv", monw32-200, 70, 10, rl.White)
	rl.DrawText(drawblockhTEXT, monw32-290, 80, 10, rl.White)
	rl.DrawText("drawblockh", monw32-200, 80, 10, rl.White)
	rl.DrawText(blocknumTEXT, monw32-290, 90, 10, rl.White)
	rl.DrawText("blocknum", monw32-200, 90, 10, rl.White)
	rl.DrawText(mouseblockTEXT, monw32-290, 100, 10, rl.White)
	rl.DrawText("mouseblock", monw32-200, 100, 10, rl.White)
	rl.DrawText(selectblockhTEXT, monw32-290, 110, 10, rl.White)
	rl.DrawText("selectblockh", monw32-200, 110, 10, rl.White)
	rl.DrawText(selectblockvTEXT, monw32-290, 120, 10, rl.White)
	rl.DrawText("selectblockv", monw32-200, 120, 10, rl.White)

}
func input() { // MARK: keys input
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if mousepos.X >= float32(menu1x) && mousepos.X < float32(menu1x+16) && mousepos.Y >= float32(menu1y) && mousepos.Y <= float32(menu1y+16) {
			if menu1moveon {
				menu1moveon = false
			} else {
				menu1moveon = true
			}
		} else if menu1moveon {
			menu1x = int32(mousepos.X)
			menu1y = int32(mousepos.Y)
			menu1moveon = false
		} else {
			playernext = mouseblock
		}
	}
	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyDown(rl.KeyRight) {
		if drawblockv < (levelw - (blocksw + 1)) {
			drawblocknext++
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyDown(rl.KeyLeft) {
		if drawblockv > 0 {
			drawblocknext--
		}
	}
	if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyDown(rl.KeyUp) {
		if drawblockh > 0 {
			drawblocknext -= levelw
		}
	}
	if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyDown(rl.KeyDown) {
		if drawblockh < levelh-(blocksh+1) {
			drawblocknext += levelw
		}
	}
	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
			zoom1 = false
			zoom4 = false
			zoom2 = true
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 4.0
			zoom1 = false
			zoom4 = true
			zoom2 = false
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 4.0 {
			camera.Zoom = 2.0
			zoom1 = false
			zoom4 = false
			zoom2 = true
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
			zoom1 = true
			zoom4 = false
			zoom2 = false
		}
	}
	if rl.IsKeyPressed(rl.KeyF1) {
		if grid16on {
			grid16on = false
		} else {
			grid16on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if grid4on {
			grid4on = false
		} else {
			grid4on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

}
func grid16() { // MARK: grid16()
	for a := 0; a < monitorw; a += 16 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.Green, 0.1))
	}
	for a := 0; a < monitorh; a += 16 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.Green, 0.1))
	}
}
func grid4() { // MARK: grid4()
	for a := 0; a < monitorw; a += 4 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.DarkGreen, 0.1))
	}
	for a := 0; a < monitorh; a += 4 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.DarkGreen, 0.1))
	}
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
