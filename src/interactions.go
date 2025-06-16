// ****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2025 MASA Group
//
// ****************************************************************************

package hla

// Service for interactions management
type Interactions interface {
	PublishInteractionClass(className string)

	SubscribeInteractionClass(className string, federate *InteractionReceiver)

	SendInteraction(className string, parametersValue map[string][]byte, time uint64)
}
