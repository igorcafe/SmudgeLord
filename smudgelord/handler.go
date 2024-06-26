package smudgelord

import (
	"smudgelord/smudgelord/database"
	"smudgelord/smudgelord/modules"
	"smudgelord/smudgelord/modules/lastfm"
	"smudgelord/smudgelord/modules/medias"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Handler struct {
	bot *telego.Bot
	bh  *th.BotHandler
}

func NewHandler(bot *telego.Bot, bh *th.BotHandler) *Handler {
	return &Handler{
		bot: bot,
		bh:  bh,
	}
}

func (h *Handler) RegisterHandlers() {
	// Add middleware
	h.bh.Use(database.SaveUsers)
	h.bh.Use(modules.CheckAFK)

	// Add module handlers
	modules.LoadStart(h.bh, h.bot)
	modules.LoadAFK(h.bh, h.bot)
	lastfm.LoadLastFM(h.bh, h.bot)
	medias.LoadMediaDownloader(h.bh, h.bot)
	modules.LoadMisc(h.bh, h.bot)
	modules.LoadStickers(h.bh, h.bot)
	modules.LoadSudoers(h.bh, h.bot)
}
