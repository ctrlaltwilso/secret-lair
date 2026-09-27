package mail

import (
	"fmt"
	"log"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

func MailClient(username string, app_password string, inbox string, address string) error {
	c, err := Connect(username, app_password, address)
	if err != nil {
		return err
	}
	defer c.Close()

	m, err := CheckMailbox(c, inbox)
	if err != nil {
		return err
	}

	err = FetchLatestMessage(m, c, inbox)
	if err != nil {
		return err
	}

	if err := c.Logout().Wait(); err != nil {
		return fmt.Errorf("failed to logout: %v", err)
	}

	return nil
}

func Connect(username string, app_password string, address string) (*imapclient.Client, error) {
	c, err := imapclient.DialTLS(address, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to dial IMAP server: %v", err)
	}

	if err := c.Login(username, app_password).Wait(); err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	return c, nil
}

func CheckMailbox(client *imapclient.Client, inbox string) (*imap.SelectData, error) {

	selectedMbox, err := client.Select(inbox, nil).Wait()
	if err != nil {
		mailboxes, listErr := ListMailBoxes(client)
		if listErr == nil {
			log.Printf("available mailboxes: %v", mailboxes)
		}
		return nil, fmt.Errorf("failed to select %s: %v", inbox, err)
	}

	log.Printf("%s contains %v messages", inbox, selectedMbox.NumMessages)
	return selectedMbox, nil

}

func ListMailBoxes(client *imapclient.Client) ([]string, error) {

	mailboxes, err := client.List("", "%", nil).Collect()
	if err != nil {
		return nil, fmt.Errorf("failed to collect mailboxes: %v", err)
	}
	names := make([]string, len(mailboxes))
	for i, mbox := range mailboxes {
		names[i] = mbox.Mailbox
	}

	return names, nil
}

func FetchLatestMessage(m *imap.SelectData, c *imapclient.Client, inbox string) error {

	if m.NumMessages > 0 {
		seqSet := imap.SeqSetNum(m.NumMessages)
		fetchOptions := &imap.FetchOptions{Envelope: true}
		messages, err := c.Fetch(seqSet, fetchOptions).Collect()
		if err != nil {
			return fmt.Errorf("failed to fetch first message in %s: %v", inbox, err)
		}

		log.Printf("subject of most recent message in %s: %v", inbox, messages[0].Envelope.Subject)

	}

	return nil
}
