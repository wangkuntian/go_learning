package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

// 并发不安全
func loadIcon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

var loadIconsOnce sync.Once

// Icon 并发安全
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
