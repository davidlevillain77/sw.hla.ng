// ****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2025 MASA Group
//
// ****************************************************************************

package hla

// Service for objects management
type Objects interface {
	PublishObjectClass(className string, attributesName []string)

	SubscribeObjectClass(className string, attributesName []string, federate *ObjectReceiver)

	ReserveObjectName(instanceName string, className string, federate *ObjectCreator)

	UpdateObject(instanceName string, attributesValue map[string][]byte, time uint64)

	RemoveObject(instanceName string, time uint64)
}

type ObjectReceiver interface {
	ObjectUpdated(instanceName string, attributesName []string, time uint64)
}
