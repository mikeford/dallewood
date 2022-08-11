package render

import (
	"github.com/fogleman/ease"
	"github.com/thediveo/enumflag"
)

type AnimationEasingMode enumflag.Flag

const (
	EasingModeLinear AnimationEasingMode = iota
	EasingModeInQuad
	EasingModeInCubic
	EasingModeInQuart
	EasingModeInQuint
	EasingModeInSine
	EasingModeInExpo
	EasingModeInCirc
	EasingModeInElastic
	EasingModeInBack
	EasingModeInBounce
	EasingModeOutQuad
	EasingModeOutCubic
	EasingModeOutQuart
	EasingModeOutQuint
	EasingModeOutSine
	EasingModeOutExpo
	EasingModeOutCirc
	EasingModeOutElastic
	EasingModeOutBack
	EasingModeOutBounce
	EasingModeInOutQuad
	EasingModeInOutCubic
	EasingModeInOutQuart
	EasingModeInOutQuint
	EasingModeInOutSine
	EasingModeInOutExpo
	EasingModeInOutCirc
	EasingModeInOutElastic
	EasingModeInOutBack
	EasingModeInOutBounce
)

var AnimationEasingModeIds = map[AnimationEasingMode][]string{
	EasingModeLinear:       {"linear"},
	EasingModeInQuad:       {"inquad"},
	EasingModeInCubic:      {"incubic"},
	EasingModeInQuart:      {"inquart"},
	EasingModeInQuint:      {"inquint"},
	EasingModeInSine:       {"insine"},
	EasingModeInExpo:       {"inexpo"},
	EasingModeInCirc:       {"incirc"},
	EasingModeInElastic:    {"inelastic"},
	EasingModeInBack:       {"inback"},
	EasingModeInBounce:     {"inbounce"},
	EasingModeOutQuad:      {"outquad"},
	EasingModeOutCubic:     {"outcubic"},
	EasingModeOutQuart:     {"outquart"},
	EasingModeOutQuint:     {"outquint"},
	EasingModeOutSine:      {"outsine"},
	EasingModeOutExpo:      {"outexpo"},
	EasingModeOutCirc:      {"outcirc"},
	EasingModeOutElastic:   {"outelastic"},
	EasingModeOutBack:      {"outback"},
	EasingModeOutBounce:    {"outbounce"},
	EasingModeInOutQuad:    {"inoutquad"},
	EasingModeInOutCubic:   {"inoutcubic"},
	EasingModeInOutQuart:   {"inoutquart"},
	EasingModeInOutQuint:   {"inoutquint"},
	EasingModeInOutSine:    {"inoutsine"},
	EasingModeInOutExpo:    {"inoutexpo"},
	EasingModeInOutCirc:    {"inoutcirc"},
	EasingModeInOutElastic: {"inoutelastic"},
	EasingModeInOutBack:    {"inoutback"},
	EasingModeInOutBounce:  {"inoutbounce"},
}

// Takes a time t between 0 and 1 and returns the corresponding
// time for the specified progress function.
func calculateNormalizedTime(t float64, easingMode AnimationEasingMode) (normalized float64) {
	switch easingMode {
	case EasingModeLinear:
		{
			normalized = ease.Linear(t)
		}
	case EasingModeInQuad:
		{
			normalized = ease.InQuad(t)
		}
	case EasingModeInCubic:
		{
			normalized = ease.InCubic(t)
		}
	case EasingModeInQuart:
		{
			normalized = ease.InQuart(t)
		}
	case EasingModeInQuint:
		{
			normalized = ease.InQuint(t)
		}
	case EasingModeInSine:
		{
			normalized = ease.InSine(t)
		}
	case EasingModeInExpo:
		{
			normalized = ease.InExpo(t)
		}
	case EasingModeInCirc:
		{
			normalized = ease.InCirc(t)
		}
	case EasingModeInElastic:
		{
			normalized = ease.InElastic(t)
		}
	case EasingModeInBack:
		{
			normalized = ease.InBack(t)
		}
	case EasingModeInBounce:
		{
			normalized = ease.InBounce(t)
		}
	case EasingModeOutQuad:
		{
			normalized = ease.OutQuad(t)
		}
	case EasingModeOutCubic:
		{
			normalized = ease.OutCubic(t)
		}
	case EasingModeOutQuart:
		{
			normalized = ease.OutQuart(t)
		}
	case EasingModeOutQuint:
		{
			normalized = ease.OutQuint(t)
		}
	case EasingModeOutSine:
		{
			normalized = ease.OutSine(t)
		}
	case EasingModeOutExpo:
		{
			normalized = ease.OutExpo(t)
		}
	case EasingModeOutCirc:
		{
			normalized = ease.OutCirc(t)
		}
	case EasingModeOutElastic:
		{
			normalized = ease.OutElastic(t)
		}
	case EasingModeOutBack:
		{
			normalized = ease.OutBack(t)
		}
	case EasingModeOutBounce:
		{
			normalized = ease.OutBounce(t)
		}
	case EasingModeInOutQuad:
		{
			normalized = ease.InOutQuad(t)
		}
	case EasingModeInOutCubic:
		{
			normalized = ease.InOutCubic(t)
		}
	case EasingModeInOutQuart:
		{
			normalized = ease.InOutQuart(t)
		}
	case EasingModeInOutQuint:
		{
			normalized = ease.InOutQuint(t)
		}
	case EasingModeInOutSine:
		{
			normalized = ease.InOutSine(t)
		}
	case EasingModeInOutExpo:
		{
			normalized = ease.InOutExpo(t)
		}
	case EasingModeInOutCirc:
		{
			normalized = ease.InOutCirc(t)
		}
	case EasingModeInOutElastic:
		{
			normalized = ease.InOutElastic(t)
		}
	case EasingModeInOutBack:
		{
			normalized = ease.InOutBack(t)
		}
	case EasingModeInOutBounce:
		{
			normalized = ease.InOutBounce(t)
		}
	}
	return normalized
}
