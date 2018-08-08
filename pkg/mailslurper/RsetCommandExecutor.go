// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailslurper

import (
	"github.com/sirupsen/logrus"
)

/*
RsetCommandExecutor process command like RSET
*/
type RsetCommandExecutor struct {
	logger *logrus.Entry
	reader *SMTPReader
	writer *SMTPWriter
}

/*
NewRsetCommandExecutor creates a new struct
*/
func NewRsetCommandExecutor(logger *logrus.Entry, reader *SMTPReader, writer *SMTPWriter) *RsetCommandExecutor {
	return &RsetCommandExecutor{
		logger: logger,
		reader: reader,
		writer: writer,
	}
}

/*
Process handles the RSET command
*/
func (e *RsetCommandExecutor) Process(streamInput string, mailItem *MailItem) error {

	e.logger.Debugf("> %s", streamInput)
	return e.writer.SendOkResponse()
}
