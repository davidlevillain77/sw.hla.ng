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
type Time interface {
	SetRegulating(isRegulating bool)

	SetConstrained(isConstrained bool)

	ModifyLookahead(lookahead uint64)

	TimeAdvanceRequest(timeInMs uint64, federate *TimeManaged)
}

type TimeManaged interface {
	TimeAdvanceGranted(timeInMs uint64)
}
