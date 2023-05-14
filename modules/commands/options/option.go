package options

type CommandOpts struct {
	Name            string
	ChatID          int64
	MessageID       int
	ChannelUsername string
}

func GetOpt(opts ...any) CommandOpts {
	var nextOpt = new(CommandOpts)
	for i, opt := range opts {
		switch v := opt.(type) {
		case string:
			if i == 0 {
				nextOpt.Name = v
			}
			if i == 3 {
				nextOpt.ChannelUsername = v
			}
		case int64:
			if i == 1 {
				nextOpt.ChatID = v
			}

		case int:
			if i == 2 {
				nextOpt.MessageID = v
			}
		default:
			// leave emtpy
		}
	}

	return *nextOpt
}
