// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dock

import (
	"path/filepath"

	"github.com/linuxdeepin/go-lib/log"
	"github.com/linuxdeepin/go-lib/xdg/basedir"
	"github.com/linuxdeepin/dde-daemon/loader"

	x "github.com/linuxdeepin/go-x11-client"
)

func init() {
	loader.Register(NewDaemon(logger))
}

var (
	logger      = log.NewLogger("daemon/dock")
	homeDir     string
	scratchDir  string
	dockManager *Manager

	globalXConn *x.Conn

	atomNetShowingDesktop       x.Atom
	atomNetClientList           x.Atom
	atomNetActiveWindow         x.Atom
	atomNetWMIcon               x.Atom
	atomNetWMName               x.Atom
	atomNetWMState              x.Atom
	atomNetWMWindowType         x.Atom
	atomXEmbedInfo              x.Atom
	atomNetFrameExtents         x.Atom
	atomGtkFrameExtents         x.Atom
	atomNetWmStateHidden        x.Atom
	atomWmWindowRole            x.Atom
	atomUTF8String              x.Atom
	atomString                  x.Atom
	atomInteger                 x.Atom
	atomFlatpakAppId            x.Atom
	atomGtkApplicationId        x.Atom
	atomNetWmWindowOpacity      x.Atom
	atomWmClientLeader          x.Atom
	atomWmCommand               x.Atom
	atomNetWmStateFocused       x.Atom //nolint
	atomNetWmWindowTypeDesktop  x.Atom
	atomNetWmActionMinimize     x.Atom
	atomWmStateDemandsAttention x.Atom
	atomNetWmStateSkipTaskbar   x.Atom
	atomNetWmStateModal         x.Atom
	atomNetWmWindowTypeDialog   x.Atom
	atomNetWmStateMaximizedVert x.Atom
	atomNetWmStateMaximizedHorz x.Atom
	atomNetWmStateAbove         x.Atom
	atomNetWmActionClose        x.Atom
	atomNetWmAllowedActions     x.Atom
	atomNetWmPid                x.Atom
	atomMotifWmHints            x.Atom

	atomAndroidUengineId   x.Atom
	atomAndroidUengineName x.Atom
)

func initDir() {
	homeDir = basedir.GetUserHomeDir()
	scratchDir = filepath.Join(basedir.GetUserConfigDir(), "dock/scratch")
	logger.Debugf("scratch dir: %q", scratchDir)
}

func initAtom() {
	atomNetShowingDesktop, _ = getAtom("_NET_SHOWING_DESKTOP")
	atomNetClientList, _ = getAtom("_NET_CLIENT_LIST")
	atomNetActiveWindow, _ = getAtom("_NET_ACTIVE_WINDOW")
	atomNetWMIcon, _ = getAtom("_NET_WM_ICON")
	atomNetWMName, _ = getAtom("_NET_WM_NAME")
	atomNetWMState, _ = getAtom("_NET_WM_STATE")
	atomNetWMWindowType, _ = getAtom("_NET_WM_WINDOW_TYPE")
	atomXEmbedInfo, _ = getAtom("_XEMBED_INFO")
	atomNetFrameExtents, _ = getAtom("_NET_FRAME_EXTENTS")
	atomGtkFrameExtents, _ = getAtom("_GTK_FRAME_EXTENTS")
	atomNetWmStateHidden, _ = getAtom("_NET_WM_STATE_HIDDEN")
	atomWmWindowRole, _ = getAtom("WM_WINDOW_ROLE")
	atomUTF8String, _ = getAtom("UTF8_STRING")
	atomString, _ = getAtom("STRING")
	atomInteger, _ = getAtom("INTEGER")
	atomFlatpakAppId, _ = getAtom("FLATPAK_APPID")
	atomGtkApplicationId, _ = getAtom("_GTK_APPLICATION_ID")
	atomNetWmWindowOpacity, _ = getAtom("_NET_WM_WINDOW_OPACITY")
	atomWmClientLeader, _ = getAtom("WM_CLIENT_LEADER")
	atomWmCommand, _ = getAtom("WM_COMMAND")
	atomNetWmStateFocused, _ = getAtom("_NET_WM_STATE_FOCUSED")
	atomNetWmWindowTypeDesktop, _ = getAtom("_NET_WM_WINDOW_TYPE_DESKTOP")
	atomNetWmActionMinimize, _ = getAtom("_NET_WM_ACTION_MINIMIZE")
	atomWmStateDemandsAttention, _ = getAtom("_NET_WM_STATE_DEMANDS_ATTENTION")
	atomNetWmStateSkipTaskbar, _ = getAtom("_NET_WM_STATE_SKIP_TASKBAR")
	atomNetWmStateModal, _ = getAtom("_NET_WM_STATE_MODAL")
	atomNetWmWindowTypeDialog, _ = getAtom("_NET_WM_WINDOW_TYPE_DIALOG")
	atomNetWmStateMaximizedVert, _ = getAtom("_NET_WM_STATE_MAXIMIZED_VERT")
	atomNetWmStateMaximizedHorz, _ = getAtom("_NET_WM_STATE_MAXIMIZED_HORZ")
	atomNetWmStateAbove, _ = getAtom("_NET_WM_STATE_ABOVE")
	atomNetWmActionClose, _ = getAtom("_NET_WM_ACTION_CLOSE")
	atomNetWmAllowedActions, _ = getAtom("_NET_WM_ALLOWED_ACTIONS")
	atomNetWmPid, _ = getAtom("_NET_WM_PID")
	atomMotifWmHints, _ = getAtom("_MOTIF_WM_HINTS")
	atomAndroidUengineId, _ = getAtom("UENGINE_TASK_ID")
	atomAndroidUengineName, _ = getAtom("UENGINE_TASK_NAME")
}
