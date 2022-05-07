package fjira

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mk5/fjira/internal/app"
	"github.com/mk5/fjira/internal/jira"
)

const (
	ActionAssigneeChange app.ActionBarAction = iota
	ActionStatusChange
	ActionEscape
	ActionYes
	ActionNo
)

var (
	BottomBarActionBarItemBold = tcell.StyleDefault.Bold(true).Foreground(tcell.ColorDarkKhaki)
	BottomBarActionBarKeyBold  = tcell.StyleDefault.Bold(true).Foreground(tcell.ColorDarkCyan).Underline(true)
	TopBarItemBold             = tcell.StyleDefault.Bold(true).Foreground(tcell.ColorDarkKhaki)
	IssueBarActionBarItemBold  = tcell.StyleDefault.Bold(true).Foreground(tcell.ColorDarkKhaki)
)

// TODO - concrete "bottomBar" struct with helper methods like "setCurrentProject(..)"
func CreateNewEmptyProjectBottomBar() *app.ActionBar {
	actionBar := app.NewActionBar(app.Bottom, app.Left)
	actionBar.AddItemWithStyles(
		MessageProjectLabel,
		app.ActionBarLabel(""),
		tcell.StyleDefault, BottomBarActionBarItemBold,
	)
	return actionBar
}

func CreateNewIssueBottomBar(issue *jira.JiraIssue) *app.ActionBar {
	actionBar := app.NewActionBar(app.Bottom, app.Left)
	actionBar.AddItemWithStyles(
		MessageIssueLabel,
		app.ActionBarLabel(issue.Key),
		tcell.StyleDefault, BottomBarActionBarItemBold,
	)
	return actionBar
}

func CreateNewSearchIssuesBottomBar(project *jira.JiraProject) *app.ActionBar {
	actionBar := app.NewActionBar(app.Bottom, app.Left)
	actionBar.AddItemWithStyles(
		MessageProjectLabel,
		app.ActionBarLabel(fmt.Sprintf("[%s]%s", project.Key, project.Name)),
		tcell.StyleDefault, BottomBarActionBarItemBold,
	)
	actionBar.AddItem(NewByStatusBarItem())
	actionBar.AddItem(NewByAssigneeBarItem())
	return actionBar
}

// TODO - refactor - create general place with navigation definition
func CreateNewSearchIssuesTopBar() *app.ActionBar {
	actionBar := app.NewActionBar(app.Top, app.Right)
	actionBar.AddItemWithStyles(
		"Status: ",
		"All",
		tcell.StyleDefault, TopBarItemBold,
	)
	actionBar.AddItemWithStyles(
		"Assignee: ",
		"All",
		tcell.StyleDefault, TopBarItemBold,
	)
	return actionBar
}

func CreateNewIssueTopBar(issue *jira.JiraIssue) *app.ActionBar {
	actionBar := app.NewActionBar(app.Top, app.Right)
	actionBar.AddItemWithStyles(
		MessageLabelReporter,
		issue.Fields.Reporter.DisplayName,
		tcell.StyleDefault,
		IssueBarActionBarItemBold,
	)
	actionBar.AddItemWithStyles(
		MessageLabelAssignee,
		issue.Fields.Assignee.DisplayName,
		tcell.StyleDefault,
		IssueBarActionBarItemBold,
	)
	actionBar.AddItemWithStyles(
		MessageTypeStatus,
		issue.Fields.Type.Name,
		tcell.StyleDefault,
		IssueBarActionBarItemBold,
	)
	actionBar.AddItemWithStyles(
		MessageLabelStatus,
		issue.Fields.Status.Name,
		tcell.StyleDefault,
		IssueBarActionBarItemBold,
	)
	return actionBar
}

func NewCancelBarItem() *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:         int(ActionEscape),
		Text1:      "ESC",
		Text2:      " - cancel",
		Text1Style: BottomBarActionBarKeyBold,
		Text2Style: tcell.StyleDefault,
		TriggerKey: tcell.KeyEscape,
	}
}

func NewStatusChangeBarItem() *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:          int(ActionStatusChange),
		Text1:       "s",
		Text2:       " - change status",
		Text1Style:  BottomBarActionBarKeyBold,
		Text2Style:  tcell.StyleDefault,
		TriggerKey:  tcell.KeyF1,
		TriggerRune: 's',
	}
}

func NewByStatusBarItem() *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:         int(ActionStatusChange),
		Text1:      "F1",
		Text2:      " - by status",
		Text1Style: BottomBarActionBarKeyBold,
		Text2Style: tcell.StyleDefault,
		TriggerKey: tcell.KeyF1,
	}
}

func NewByAssigneeBarItem() *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:         int(ActionAssigneeChange),
		Text1:      "F2",
		Text2:      " - by assignee",
		Text1Style: BottomBarActionBarKeyBold,
		Text2Style: tcell.StyleDefault,
		TriggerKey: tcell.KeyF2,
	}
}

func NewAssigneeChangeBarItem() *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:          int(ActionAssigneeChange),
		Text1:       "a",
		Text2:       " - assign user",
		Text1Style:  BottomBarActionBarKeyBold,
		Text2Style:  tcell.StyleDefault,
		TriggerKey:  tcell.KeyF2,
		TriggerRune: 'a',
	}
}

func NewNewStatusBarItem(newStatus string) *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:         -1,
		Text1:      "New status: ",
		Text2:      newStatus,
		Text1Style: tcell.StyleDefault,
		Text2Style: BottomBarActionBarKeyBold,
	}
}

func NewNewAssigneeBarItem(newAssignee *jira.JiraUser) *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:         -1,
		Text1:      "New assignee: ",
		Text2:      fmt.Sprintf("%s <%s>", newAssignee.DisplayName, newAssignee.EmailAddress),
		Text1Style: tcell.StyleDefault,
		Text2Style: BottomBarActionBarKeyBold,
	}
}

func NewYesBarItem() *app.ActionBarItem {
	return &app.ActionBarItem{
		Id:          int(ActionYes),
		Text1:       "y",
		Text2:       " - yes",
		Text1Style:  BottomBarActionBarKeyBold,
		Text2Style:  tcell.StyleDefault,
		TriggerRune: 'y',
	}
}