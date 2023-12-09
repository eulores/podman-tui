package cntdialogs

import (
	"fmt"
	"strings"

	"github.com/containers/podman-tui/pdcs/containers"
	"github.com/containers/podman-tui/pdcs/images"
	"github.com/containers/podman-tui/pdcs/networks"
	"github.com/containers/podman-tui/pdcs/pods"
	"github.com/containers/podman-tui/pdcs/volumes"
	"github.com/containers/podman-tui/ui/dialogs"
	"github.com/containers/podman-tui/ui/style"
	"github.com/containers/podman-tui/ui/utils"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
)

const (
	containerCreateDialogMaxWidth = 100
	containerCreateDialogHeight   = 18
)

const (
	createContainerFormFocus = 0 + iota
	createCategoriesFocus
	createCategoryPagesFocus
	createContainerNameFieldFocus
	createContainerImageFieldFocus
	createcontainerPodFieldFocis
	createContainerLabelsFieldFocus
	createContainerRemoveFieldFocus
	createContainerPrivilegedFieldFocus
	createContainerTimeoutFieldFocus
	createContainerEnvHostFieldFocus
	createContainerEnvVarsFieldFocus
	createContainerEnvFileFieldFocus
	createContainerEnvMergeFieldFocus
	createContainerWorkDirFieldFocus
	createContainerUmaskFieldFocus
	createContainerUnsetEnvFieldFocus
	createContainerUnsetEnvAllFieldFocus
	createContainerUserFieldFocus
	createContainerHostUsersFieldFocus
	createContainerPasswdEntryFieldFocus
	createContainerGroupEntryFieldFocus
	createcontainerSecLabelFieldFocus
	createContainerApprarmorFieldFocus
	createContainerSeccompFeildFocus
	createcontainerSecMaskFieldFocus
	createcontainerSecUnmaskFieldFocus
	createcontainerSecNoNewPrivFieldFocus
	createContainerPortExposeFieldFocus
	createContainerPortPublishFieldFocus
	createContainerPortPublishAllFieldFocus
	createContainerHostnameFieldFocus
	createContainerIPAddrFieldFocus
	createContainerMacAddrFieldFocus
	createContainerNetworkFieldFocus
	createContainerDNSServersFieldFocus
	createContainerDNSOptionsFieldFocus
	createContainerDNSSearchFieldFocus
	createContainerImageVolumeFieldFocus
	createContainerVolumeFieldFocus
	createContainerMountFieldFocus
	containerHealthCmdFieldFocus
	containerHealthStartupCmdFieldFocus
	containerHealthOnFailureFieldFocus
	containerHealthIntervalFieldFocus
	containerHealthStartupIntervalFieldFocus
	containerHealthTimeoutFieldFocus
	containerHealthStartupTimeoutFieldFocus
	containerHealthRetriesFieldFocus
	containerHealthStartupRetriesFieldFocus
	containerHealthStartPeriodFieldFocus
	containerHealthStartupSuccessFieldFocus
)

const (
	containerInfoPageIndex = 0 + iota
	environmentPageIndex
	userGroupsPageIndex
	dnsPageIndex
	healthPageIndex
	networkingPageIndex
	portPageIndex
	securityOptsPageIndex
	volumePageIndex
)

// ContainerCreateDialog implements container create dialog.
type ContainerCreateDialog struct {
	*tview.Box
	layout                              *tview.Flex
	categoryLabels                      []string
	categories                          *tview.TextView
	categoryPages                       *tview.Pages
	containerInfoPage                   *tview.Flex
	environmentPage                     *tview.Flex
	userGroupsPage                      *tview.Flex
	securityOptsPage                    *tview.Flex
	portPage                            *tview.Flex
	networkingPage                      *tview.Flex
	dnsPage                             *tview.Flex
	volumePage                          *tview.Flex
	healthPage                          *tview.Flex
	form                                *tview.Form
	display                             bool
	activePageIndex                     int
	focusElement                        int
	imageList                           []images.ImageListReporter
	podList                             []*entities.ListPodsReport
	containerNameField                  *tview.InputField
	containerImageField                 *tview.DropDown
	containerPodField                   *tview.DropDown
	containerLabelsField                *tview.InputField
	containerRemoveField                *tview.Checkbox
	containerPrivilegedField            *tview.Checkbox
	containerTimeoutField               *tview.InputField
	containerWorkDirField               *tview.InputField
	containerEnvHostField               *tview.Checkbox
	containerEnvVarsField               *tview.InputField
	containerEnvFileField               *tview.InputField
	containerEnvMergeField              *tview.InputField
	containerUmaskField                 *tview.InputField
	containerUnsetEnvField              *tview.InputField
	containerUnsetEnvAllField           *tview.Checkbox
	containerUserField                  *tview.InputField
	containerHostUsersField             *tview.InputField
	containerPasswdEntryField           *tview.InputField
	containerGroupEntryField            *tview.InputField
	containerSecLabelField              *tview.InputField
	containerSecApparmorField           *tview.InputField
	containerSeccompField               *tview.InputField
	containerSecMaskField               *tview.InputField
	containerSecUnmaskField             *tview.InputField
	containerSecNoNewPrivField          *tview.Checkbox
	containerPortExposeField            *tview.InputField
	containerPortPublishField           *tview.InputField
	ContainerPortPublishAllField        *tview.Checkbox
	containerHostnameField              *tview.InputField
	containerIPAddrField                *tview.InputField
	containerMacAddrField               *tview.InputField
	containerNetworkField               *tview.DropDown
	containerDNSServersField            *tview.InputField
	containerDNSOptionsField            *tview.InputField
	containerDNSSearchField             *tview.InputField
	containerHealthCmdField             *tview.InputField
	containerHealthIntervalField        *tview.InputField
	containerHealthOnFailureField       *tview.DropDown
	containerHealthRetriesField         *tview.InputField
	containerHealthStartPeriodField     *tview.InputField
	containerHealthTimeoutField         *tview.InputField
	containerHealthStartupCmdField      *tview.InputField
	containerHealthStartupIntervalField *tview.InputField
	containerHealthStartupRetriesField  *tview.InputField
	containerHealthStartupSuccessField  *tview.InputField
	containerHealthStartupTimeoutField  *tview.InputField
	containerVolumeField                *tview.InputField
	containerImageVolumeField           *tview.DropDown
	containerMountField                 *tview.InputField
	cancelHandler                       func()
	createHandler                       func()
}

// NewContainerCreateDialog returns new container create dialog primitive ContainerCreateDialog.
func NewContainerCreateDialog() *ContainerCreateDialog {
	containerDialog := ContainerCreateDialog{
		Box:               tview.NewBox(),
		layout:            tview.NewFlex().SetDirection(tview.FlexRow),
		categories:        tview.NewTextView(),
		categoryPages:     tview.NewPages(),
		containerInfoPage: tview.NewFlex(),
		environmentPage:   tview.NewFlex(),
		userGroupsPage:    tview.NewFlex(),
		securityOptsPage:  tview.NewFlex(),
		networkingPage:    tview.NewFlex(),
		dnsPage:           tview.NewFlex(),
		portPage:          tview.NewFlex(),
		volumePage:        tview.NewFlex(),
		healthPage:        tview.NewFlex(),
		form:              tview.NewForm(),
		categoryLabels: []string{
			"Container",
			"Environment",
			"User and groups",
			"DNS Settings",
			"Health check",
			"Network Settings",
			"Ports Settings",
			"Security Options",
			"Volumes Settings",
		},
		activePageIndex:                     0,
		display:                             false,
		containerNameField:                  tview.NewInputField(),
		containerImageField:                 tview.NewDropDown(),
		containerPodField:                   tview.NewDropDown(),
		containerLabelsField:                tview.NewInputField(),
		containerRemoveField:                tview.NewCheckbox(),
		containerPrivilegedField:            tview.NewCheckbox(),
		containerTimeoutField:               tview.NewInputField(),
		containerWorkDirField:               tview.NewInputField(),
		containerEnvHostField:               tview.NewCheckbox(),
		containerEnvVarsField:               tview.NewInputField(),
		containerEnvFileField:               tview.NewInputField(),
		containerEnvMergeField:              tview.NewInputField(),
		containerUmaskField:                 tview.NewInputField(),
		containerUnsetEnvField:              tview.NewInputField(),
		containerUnsetEnvAllField:           tview.NewCheckbox(),
		containerUserField:                  tview.NewInputField(),
		containerHostUsersField:             tview.NewInputField(),
		containerPasswdEntryField:           tview.NewInputField(),
		containerGroupEntryField:            tview.NewInputField(),
		containerSecLabelField:              tview.NewInputField(),
		containerSecApparmorField:           tview.NewInputField(),
		containerSeccompField:               tview.NewInputField(),
		containerSecMaskField:               tview.NewInputField(),
		containerSecUnmaskField:             tview.NewInputField(),
		containerSecNoNewPrivField:          tview.NewCheckbox(),
		containerPortExposeField:            tview.NewInputField(),
		containerPortPublishField:           tview.NewInputField(),
		ContainerPortPublishAllField:        tview.NewCheckbox(),
		containerHostnameField:              tview.NewInputField(),
		containerIPAddrField:                tview.NewInputField(),
		containerMacAddrField:               tview.NewInputField(),
		containerNetworkField:               tview.NewDropDown(),
		containerDNSServersField:            tview.NewInputField(),
		containerDNSOptionsField:            tview.NewInputField(),
		containerDNSSearchField:             tview.NewInputField(),
		containerVolumeField:                tview.NewInputField(),
		containerImageVolumeField:           tview.NewDropDown(),
		containerMountField:                 tview.NewInputField(),
		containerHealthCmdField:             tview.NewInputField(),
		containerHealthIntervalField:        tview.NewInputField(),
		containerHealthOnFailureField:       tview.NewDropDown(),
		containerHealthRetriesField:         tview.NewInputField(),
		containerHealthStartPeriodField:     tview.NewInputField(),
		containerHealthTimeoutField:         tview.NewInputField(),
		containerHealthStartupCmdField:      tview.NewInputField(),
		containerHealthStartupIntervalField: tview.NewInputField(),
		containerHealthStartupRetriesField:  tview.NewInputField(),
		containerHealthStartupSuccessField:  tview.NewInputField(),
		containerHealthStartupTimeoutField:  tview.NewInputField(),
	}

	containerDialog.setupLayout()
	containerDialog.setActiveCategory(0)
	containerDialog.initCustomInputHanlers()

	return &containerDialog
}

func (d *ContainerCreateDialog) setupLayout() {
	bgColor := style.DialogBgColor

	d.categories.SetDynamicColors(true).
		SetWrap(true).
		SetTextAlign(tview.AlignLeft)
	d.categories.SetBackgroundColor(bgColor)
	d.categories.SetBorder(true)
	d.categories.SetBorderColor(style.DialogSubBoxBorderColor)

	// category pages
	d.categoryPages.SetBackgroundColor(bgColor)
	d.categoryPages.SetBorder(true)
	d.categoryPages.SetBorderColor(style.DialogSubBoxBorderColor)

	d.setupContainerInfoPageUI()
	d.setupEnvironmentPageUI()
	d.setupUserGroupsPageUI()
	d.setupDNSPageUI()
	d.setupHealthPageUI()
	d.setupNetworkPageUI()
	d.setupPortsPageUI()
	d.setupSecurityPageUI()
	d.setupVolumePageUI()

	// form
	d.form.SetBackgroundColor(bgColor)
	d.form.AddButton("Cancel", nil)
	d.form.AddButton("Create", nil)
	d.form.SetButtonsAlign(tview.AlignRight)
	d.form.SetButtonBackgroundColor(style.ButtonBgColor)

	// adding category pages
	d.categoryPages.AddPage(d.categoryLabels[containerInfoPageIndex], d.containerInfoPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[environmentPageIndex], d.environmentPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[userGroupsPageIndex], d.userGroupsPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[dnsPageIndex], d.dnsPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[healthPageIndex], d.healthPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[networkingPageIndex], d.networkingPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[portPageIndex], d.portPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[securityOptsPageIndex], d.securityOptsPage, true, true)
	d.categoryPages.AddPage(d.categoryLabels[volumePageIndex], d.volumePage, true, true)

	// add it to layout.
	d.layout.SetBackgroundColor(bgColor)
	d.layout.SetBorder(true)
	d.layout.SetBorderColor(style.DialogBorderColor)
	d.layout.SetTitle("PODMAN CONTAINER CREATE")

	_, layoutWidth := utils.AlignStringListWidth(d.categoryLabels)
	layout := tview.NewFlex().SetDirection(tview.FlexColumn)

	layout.AddItem(d.categories, layoutWidth+6, 0, true) //nolint:gomnd
	layout.AddItem(d.categoryPages, 0, 1, true)
	layout.SetBackgroundColor(bgColor)
	d.layout.AddItem(layout, 0, 1, true)

	d.layout.AddItem(d.form, dialogs.DialogFormHeight, 0, true)
}

func (d *ContainerCreateDialog) setupContainerInfoPageUI() {
	bgColor := style.DialogBgColor
	ddUnselectedStyle := style.DropDownUnselected
	ddselectedStyle := style.DropDownSelected
	inputFieldBgColor := style.InputFieldBgColor
	cntInfoPageLabelWidth := 12

	// name field
	d.containerNameField.SetLabel("name:")
	d.containerNameField.SetLabelWidth(cntInfoPageLabelWidth)
	d.containerNameField.SetBackgroundColor(bgColor)
	d.containerNameField.SetLabelColor(style.DialogFgColor)
	d.containerNameField.SetFieldBackgroundColor(inputFieldBgColor)

	// image field
	d.containerImageField.SetLabel("image:")
	d.containerImageField.SetLabelWidth(cntInfoPageLabelWidth)
	d.containerImageField.SetBackgroundColor(bgColor)
	d.containerImageField.SetLabelColor(style.DialogFgColor)
	d.containerImageField.SetListStyles(ddUnselectedStyle, ddselectedStyle)
	d.containerImageField.SetFieldBackgroundColor(inputFieldBgColor)

	// pod field
	d.containerPodField.SetLabel("pod:")
	d.containerPodField.SetLabelWidth(cntInfoPageLabelWidth)
	d.containerPodField.SetBackgroundColor(bgColor)
	d.containerPodField.SetLabelColor(style.DialogFgColor)
	d.containerPodField.SetListStyles(ddUnselectedStyle, ddselectedStyle)
	d.containerPodField.SetFieldBackgroundColor(inputFieldBgColor)

	// labels field
	d.containerLabelsField.SetLabel("labels:")
	d.containerLabelsField.SetLabelWidth(cntInfoPageLabelWidth)
	d.containerLabelsField.SetBackgroundColor(bgColor)
	d.containerLabelsField.SetLabelColor(style.DialogFgColor)
	d.containerLabelsField.SetFieldBackgroundColor(inputFieldBgColor)

	// privileged
	d.containerPrivilegedField.SetLabel("privileged:")
	d.containerPrivilegedField.SetLabelWidth(cntInfoPageLabelWidth)
	d.containerPrivilegedField.SetBackgroundColor(bgColor)
	d.containerPrivilegedField.SetLabelColor(style.DialogFgColor)
	d.containerPrivilegedField.SetFieldBackgroundColor(inputFieldBgColor)

	// remove field
	removeLabel := "remove:"

	d.containerRemoveField.SetLabel(removeLabel)
	d.containerRemoveField.SetLabelWidth(len(removeLabel) + 1)
	d.containerRemoveField.SetBackgroundColor(bgColor)
	d.containerRemoveField.SetLabelColor(style.DialogFgColor)
	d.containerRemoveField.SetFieldBackgroundColor(inputFieldBgColor)

	// timeout field
	timeoutLabel := "timeout:"

	d.containerTimeoutField.SetLabel(timeoutLabel)
	d.containerTimeoutField.SetLabelWidth(len(timeoutLabel) + 1)
	d.containerTimeoutField.SetBackgroundColor(bgColor)
	d.containerTimeoutField.SetLabelColor(style.DialogFgColor)
	d.containerTimeoutField.SetFieldBackgroundColor(inputFieldBgColor)

	// layout
	labelPaddings := 4
	checkBoxLayout := tview.NewFlex().SetDirection(tview.FlexColumn)

	checkBoxLayout.SetBackgroundColor(bgColor)
	checkBoxLayout.AddItem(d.containerPrivilegedField, cntInfoPageLabelWidth+labelPaddings, 0, false)
	checkBoxLayout.AddItem(d.containerRemoveField, len(removeLabel)+labelPaddings, 0, false)
	checkBoxLayout.AddItem(d.containerTimeoutField, 0, 1, false)
	checkBoxLayout.AddItem(utils.EmptyBoxSpace(bgColor), 0, 1, true)

	d.containerInfoPage.SetDirection(tview.FlexRow)
	d.containerInfoPage.AddItem(d.containerNameField, 1, 0, true)
	d.containerInfoPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.containerInfoPage.AddItem(d.containerImageField, 1, 0, true)
	d.containerInfoPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.containerInfoPage.AddItem(d.containerPodField, 1, 0, true)
	d.containerInfoPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.containerInfoPage.AddItem(d.containerLabelsField, 1, 0, true)
	d.containerInfoPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.containerInfoPage.AddItem(checkBoxLayout, 1, 0, true)
	d.containerInfoPage.SetBackgroundColor(bgColor)
}

func (d *ContainerCreateDialog) setupEnvironmentPageUI() {
	bgColor := style.DialogBgColor
	inputFieldBgColor := style.InputFieldBgColor
	envPageLabelWidth := 12

	// environment host
	d.containerEnvHostField.SetLabel("env host:")
	d.containerEnvHostField.SetLabelWidth(envPageLabelWidth)
	d.containerEnvHostField.SetBackgroundColor(bgColor)
	d.containerEnvHostField.SetLabelColor(style.DialogFgColor)
	d.containerEnvHostField.SetFieldBackgroundColor(inputFieldBgColor)

	// unset all
	unsetEnvAllLabel := "unsetenv all"
	d.containerUnsetEnvAllField.SetLabel(unsetEnvAllLabel)
	d.containerUnsetEnvAllField.SetLabelWidth(len(unsetEnvAllLabel) + 1)
	d.containerUnsetEnvAllField.SetBackgroundColor(bgColor)
	d.containerUnsetEnvAllField.SetLabelColor(style.DialogFgColor)
	d.containerUnsetEnvAllField.SetFieldBackgroundColor(inputFieldBgColor)

	// environment variables
	d.containerEnvVarsField.SetLabel("env vars:")
	d.containerEnvVarsField.SetLabelWidth(envPageLabelWidth)
	d.containerEnvVarsField.SetBackgroundColor(bgColor)
	d.containerEnvVarsField.SetLabelColor(style.DialogFgColor)
	d.containerEnvVarsField.SetFieldBackgroundColor(inputFieldBgColor)

	// environment variables file
	d.containerEnvFileField.SetLabel("env file:")
	d.containerEnvFileField.SetLabelWidth(envPageLabelWidth)
	d.containerEnvFileField.SetBackgroundColor(bgColor)
	d.containerEnvFileField.SetLabelColor(style.DialogFgColor)
	d.containerEnvFileField.SetFieldBackgroundColor(inputFieldBgColor)

	// environment merge
	d.containerEnvMergeField.SetLabel("env merge:")
	d.containerEnvMergeField.SetLabelWidth(envPageLabelWidth)
	d.containerEnvMergeField.SetBackgroundColor(bgColor)
	d.containerEnvMergeField.SetLabelColor(style.DialogFgColor)
	d.containerEnvMergeField.SetFieldBackgroundColor(inputFieldBgColor)

	// environment unset variables
	d.containerUnsetEnvField.SetLabel("unset env:")
	d.containerUnsetEnvField.SetLabelWidth(envPageLabelWidth)
	d.containerUnsetEnvField.SetBackgroundColor(bgColor)
	d.containerUnsetEnvField.SetLabelColor(style.DialogFgColor)
	d.containerUnsetEnvField.SetFieldBackgroundColor(inputFieldBgColor)

	// working directory
	d.containerWorkDirField.SetLabel("work dir:")
	d.containerWorkDirField.SetLabelWidth(envPageLabelWidth)
	d.containerWorkDirField.SetBackgroundColor(bgColor)
	d.containerWorkDirField.SetLabelColor(style.DialogFgColor)
	d.containerWorkDirField.SetFieldBackgroundColor(inputFieldBgColor)

	// umask
	umaskLabel := "umask:"
	d.containerUmaskField.SetLabel(umaskLabel)
	d.containerUmaskField.SetLabelWidth(len(umaskLabel) + 1)
	d.containerUmaskField.SetBackgroundColor(bgColor)
	d.containerUmaskField.SetLabelColor(style.DialogFgColor)
	d.containerUmaskField.SetFieldBackgroundColor(inputFieldBgColor)

	// layout
	labelPaddings := 4
	checkBoxLayout := tview.NewFlex().SetDirection(tview.FlexColumn)

	checkBoxLayout.SetBackgroundColor(bgColor)
	checkBoxLayout.AddItem(d.containerEnvHostField, envPageLabelWidth+labelPaddings, 0, false)
	checkBoxLayout.AddItem(d.containerUnsetEnvAllField, len(unsetEnvAllLabel)+labelPaddings, 0, false)
	checkBoxLayout.AddItem(d.containerUmaskField, 0, 1, false)
	checkBoxLayout.AddItem(utils.EmptyBoxSpace(bgColor), 0, 1, true)

	d.environmentPage.SetDirection(tview.FlexRow)
	d.environmentPage.AddItem(d.containerWorkDirField, 1, 0, true)
	d.environmentPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.environmentPage.AddItem(d.containerEnvVarsField, 1, 0, true)
	d.environmentPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.environmentPage.AddItem(d.containerEnvFileField, 1, 0, true)
	d.environmentPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.environmentPage.AddItem(d.containerEnvMergeField, 1, 0, true)
	d.environmentPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.environmentPage.AddItem(d.containerUnsetEnvField, 1, 0, true)
	d.environmentPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.environmentPage.AddItem(checkBoxLayout, 1, 0, true)
	d.environmentPage.SetBackgroundColor(bgColor)
}

func (d *ContainerCreateDialog) setupUserGroupsPageUI() {
	bgColor := style.DialogBgColor
	inputFieldBgColor := style.InputFieldBgColor
	userGroupLabelWidth := 14
	userFieldWidth := 30

	// user
	d.containerUserField.SetLabel("user:")
	d.containerUserField.SetLabelWidth(userGroupLabelWidth)
	d.containerUserField.SetBackgroundColor(bgColor)
	d.containerUserField.SetLabelColor(style.DialogFgColor)
	d.containerUserField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerUserField.SetFieldWidth(userFieldWidth)

	// host users
	d.containerHostUsersField.SetLabel("host user:")
	d.containerHostUsersField.SetLabelWidth(userGroupLabelWidth)
	d.containerHostUsersField.SetBackgroundColor(bgColor)
	d.containerHostUsersField.SetLabelColor(style.DialogFgColor)
	d.containerHostUsersField.SetFieldBackgroundColor(inputFieldBgColor)

	// passwd entry
	d.containerPasswdEntryField.SetLabel("passwd entry:")
	d.containerPasswdEntryField.SetLabelWidth(userGroupLabelWidth)
	d.containerPasswdEntryField.SetBackgroundColor(bgColor)
	d.containerPasswdEntryField.SetLabelColor(style.DialogFgColor)
	d.containerPasswdEntryField.SetFieldBackgroundColor(inputFieldBgColor)

	// group entry
	d.containerGroupEntryField.SetLabel("group entry:")
	d.containerGroupEntryField.SetLabelWidth(userGroupLabelWidth)
	d.containerGroupEntryField.SetBackgroundColor(bgColor)
	d.containerGroupEntryField.SetLabelColor(style.DialogFgColor)
	d.containerGroupEntryField.SetFieldBackgroundColor(inputFieldBgColor)

	d.userGroupsPage.SetDirection(tview.FlexRow)
	d.userGroupsPage.AddItem(d.containerUserField, 1, 0, true)
	d.userGroupsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.userGroupsPage.AddItem(d.containerHostUsersField, 1, 0, true)
	d.userGroupsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.userGroupsPage.AddItem(d.containerPasswdEntryField, 1, 0, true)
	d.userGroupsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.userGroupsPage.AddItem(d.containerGroupEntryField, 1, 0, true)
	d.userGroupsPage.SetBackgroundColor(bgColor)
}

func (d *ContainerCreateDialog) setupDNSPageUI() {
	bgColor := style.DialogBgColor
	inputFieldBgColor := style.InputFieldBgColor
	dnsPageLabelWidth := 13

	// hostname field
	d.containerDNSServersField.SetLabel("dns servers:")
	d.containerDNSServersField.SetLabelWidth(dnsPageLabelWidth)
	d.containerDNSServersField.SetBackgroundColor(bgColor)
	d.containerDNSServersField.SetLabelColor(style.DialogFgColor)
	d.containerDNSServersField.SetFieldBackgroundColor(inputFieldBgColor)

	// IP field
	d.containerDNSOptionsField.SetLabel("dns options:")
	d.containerDNSOptionsField.SetLabelWidth(dnsPageLabelWidth)
	d.containerDNSOptionsField.SetBackgroundColor(bgColor)
	d.containerDNSOptionsField.SetLabelColor(style.DialogFgColor)
	d.containerDNSOptionsField.SetFieldBackgroundColor(inputFieldBgColor)

	// mac field
	d.containerDNSSearchField.SetLabel("dns search:")
	d.containerDNSSearchField.SetLabelWidth(dnsPageLabelWidth)
	d.containerDNSSearchField.SetBackgroundColor(bgColor)
	d.containerDNSSearchField.SetLabelColor(style.DialogFgColor)
	d.containerDNSSearchField.SetFieldBackgroundColor(inputFieldBgColor)

	d.dnsPage.SetDirection(tview.FlexRow)
	d.dnsPage.AddItem(d.containerDNSServersField, 1, 0, true)
	d.dnsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.dnsPage.AddItem(d.containerDNSOptionsField, 1, 0, true)
	d.dnsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.dnsPage.AddItem(d.containerDNSSearchField, 1, 0, true)
	d.dnsPage.SetBackgroundColor(bgColor)
}

func (d *ContainerCreateDialog) setupHealthPageUI() {
	bgColor := style.DialogBgColor
	inputFieldBgColor := style.InputFieldBgColor
	ddUnselectedStyle := style.DropDownUnselected
	ddselectedStyle := style.DropDownSelected
	healthPageLabelWidth := 13
	healthPageSecColLabelWidth := 18
	healthPageMultiRowFieldWidth := 7

	// health cmd
	d.containerHealthCmdField.SetLabel("Command:")
	d.containerHealthCmdField.SetLabelWidth(healthPageLabelWidth)
	d.containerHealthCmdField.SetBackgroundColor(bgColor)
	d.containerHealthCmdField.SetLabelColor(style.DialogFgColor)
	d.containerHealthCmdField.SetFieldBackgroundColor(inputFieldBgColor)

	// startup cmd
	d.containerHealthStartupCmdField.SetLabel("Startup cmd:")
	d.containerHealthStartupCmdField.SetLabelWidth(healthPageLabelWidth)
	d.containerHealthStartupCmdField.SetBackgroundColor(bgColor)
	d.containerHealthStartupCmdField.SetLabelColor(style.DialogFgColor)
	d.containerHealthStartupCmdField.SetFieldBackgroundColor(inputFieldBgColor)

	// multi primitive row01
	// startup success
	d.containerHealthStartupSuccessField.SetLabel("Startup success:")
	d.containerHealthStartupSuccessField.SetLabelWidth(healthPageSecColLabelWidth)
	d.containerHealthStartupSuccessField.SetBackgroundColor(bgColor)
	d.containerHealthStartupSuccessField.SetLabelColor(style.DialogFgColor)
	d.containerHealthStartupSuccessField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthStartupSuccessField.SetFieldWidth(healthPageMultiRowFieldWidth)

	// on-failure
	d.containerHealthOnFailureField.SetOptions([]string{"none", "kill", "restart", "stop"}, nil)
	d.containerHealthOnFailureField.SetLabel("On failure:")
	d.containerHealthOnFailureField.SetLabelWidth(healthPageLabelWidth)
	d.containerHealthOnFailureField.SetBackgroundColor(bgColor)
	d.containerHealthOnFailureField.SetLabelColor(style.DialogFgColor)
	d.containerHealthOnFailureField.SetListStyles(ddUnselectedStyle, ddselectedStyle)
	d.containerHealthOnFailureField.SetFieldBackgroundColor(inputFieldBgColor)

	// start period
	startPeroidLabel := fmt.Sprintf("%15s: ", "Start period")
	d.containerHealthStartPeriodField.SetLabel(startPeroidLabel)
	d.containerHealthStartPeriodField.SetBackgroundColor(bgColor)
	d.containerHealthStartPeriodField.SetLabelColor(style.DialogFgColor)
	d.containerHealthStartPeriodField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthStartPeriodField.SetFieldWidth(healthPageMultiRowFieldWidth)

	multiItemRow01 := tview.NewFlex().SetDirection(tview.FlexColumn)
	multiItemRow01.AddItem(d.containerHealthOnFailureField, 0, 1, true)
	multiItemRow01.AddItem(d.containerHealthStartupSuccessField, 0, 1, true)
	multiItemRow01.AddItem(d.containerHealthStartPeriodField, 0, 1, true)
	multiItemRow01.SetBackgroundColor(bgColor)

	// multi primitive row02
	// startup interval
	d.containerHealthStartupIntervalField.SetLabel("Startup interval:")
	d.containerHealthStartupIntervalField.SetLabelWidth(healthPageSecColLabelWidth)
	d.containerHealthStartupIntervalField.SetBackgroundColor(bgColor)
	d.containerHealthStartupIntervalField.SetLabelColor(style.DialogFgColor)
	d.containerHealthStartupIntervalField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthStartupIntervalField.SetFieldWidth(healthPageMultiRowFieldWidth)

	// interval
	d.containerHealthIntervalField.SetLabel("Interval:")
	d.containerHealthIntervalField.SetLabelWidth(healthPageLabelWidth)
	d.containerHealthIntervalField.SetBackgroundColor(bgColor)
	d.containerHealthIntervalField.SetLabelColor(style.DialogFgColor)
	d.containerHealthIntervalField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthIntervalField.SetFieldWidth(healthPageMultiRowFieldWidth)

	multiItemRow02 := tview.NewFlex().SetDirection(tview.FlexColumn)
	multiItemRow02.AddItem(d.containerHealthIntervalField, 0, 1, true)
	multiItemRow02.AddItem(d.containerHealthStartupIntervalField, 0, 1, true)
	multiItemRow02.AddItem(utils.EmptyBoxSpace(bgColor), 0, 1, true)
	multiItemRow02.SetBackgroundColor(bgColor)

	// multi primitive row03
	// startup retries
	d.containerHealthStartupRetriesField.SetLabel("Startup retries:")
	d.containerHealthStartupRetriesField.SetLabelWidth(healthPageSecColLabelWidth)
	d.containerHealthStartupRetriesField.SetBackgroundColor(bgColor)
	d.containerHealthStartupRetriesField.SetLabelColor(style.DialogFgColor)
	d.containerHealthStartupRetriesField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthStartupRetriesField.SetFieldWidth(healthPageMultiRowFieldWidth)

	// retires
	d.containerHealthRetriesField.SetLabel("Retries:")
	d.containerHealthRetriesField.SetLabelWidth(healthPageLabelWidth)
	d.containerHealthRetriesField.SetBackgroundColor(bgColor)
	d.containerHealthRetriesField.SetLabelColor(style.DialogFgColor)
	d.containerHealthRetriesField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthRetriesField.SetFieldWidth(healthPageMultiRowFieldWidth)

	multiItemRow03 := tview.NewFlex().SetDirection(tview.FlexColumn)
	multiItemRow03.AddItem(d.containerHealthRetriesField, 0, 1, true)
	multiItemRow03.AddItem(d.containerHealthStartupRetriesField, 0, 1, true)
	multiItemRow03.AddItem(utils.EmptyBoxSpace(bgColor), 0, 1, true)
	multiItemRow03.SetBackgroundColor(bgColor)

	// multi primitive row04
	// startup timeout
	d.containerHealthStartupTimeoutField.SetLabel("Startup timeout:")
	d.containerHealthStartupTimeoutField.SetLabelWidth(healthPageSecColLabelWidth)
	d.containerHealthStartupTimeoutField.SetBackgroundColor(bgColor)
	d.containerHealthStartupTimeoutField.SetLabelColor(style.DialogFgColor)
	d.containerHealthStartupTimeoutField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthStartupTimeoutField.SetFieldWidth(healthPageMultiRowFieldWidth)

	// timeout
	d.containerHealthTimeoutField.SetLabel("Timeout:")
	d.containerHealthTimeoutField.SetLabelWidth(healthPageLabelWidth)
	d.containerHealthTimeoutField.SetBackgroundColor(bgColor)
	d.containerHealthTimeoutField.SetLabelColor(style.DialogFgColor)
	d.containerHealthTimeoutField.SetFieldBackgroundColor(inputFieldBgColor)
	d.containerHealthTimeoutField.SetFieldWidth(healthPageMultiRowFieldWidth)

	multiItemRow04 := tview.NewFlex().SetDirection(tview.FlexColumn)
	multiItemRow04.AddItem(d.containerHealthTimeoutField, 0, 1, true)
	multiItemRow04.AddItem(d.containerHealthStartupTimeoutField, 0, 1, true)
	multiItemRow04.AddItem(utils.EmptyBoxSpace(bgColor), 0, 1, true)
	multiItemRow04.SetBackgroundColor(bgColor)

	// health page
	d.healthPage.SetDirection(tview.FlexRow)
	d.healthPage.AddItem(d.containerHealthCmdField, 1, 0, true)
	d.healthPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.healthPage.AddItem(d.containerHealthStartupCmdField, 1, 0, true)
	d.healthPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.healthPage.AddItem(multiItemRow01, 1, 0, true)
	d.healthPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.healthPage.AddItem(multiItemRow02, 1, 0, true)
	d.healthPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.healthPage.AddItem(multiItemRow03, 1, 0, true)
	d.healthPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.healthPage.AddItem(multiItemRow04, 1, 0, true)
	d.healthPage.SetBackgroundColor(bgColor)
}

func (d *ContainerCreateDialog) setupNetworkPageUI() {
	bgColor := style.DialogBgColor
	ddUnselectedStyle := style.DropDownUnselected
	ddselectedStyle := style.DropDownSelected
	inputFieldBgColor := style.InputFieldBgColor
	networkingPageLabelWidth := 13

	// hostname field
	d.containerHostnameField.SetLabel("hostname:")
	d.containerHostnameField.SetLabelWidth(networkingPageLabelWidth)
	d.containerHostnameField.SetBackgroundColor(bgColor)
	d.containerHostnameField.SetLabelColor(style.DialogFgColor)
	d.containerHostnameField.SetFieldBackgroundColor(inputFieldBgColor)

	// IP field
	d.containerIPAddrField.SetLabel("ip address:")
	d.containerIPAddrField.SetLabelWidth(networkingPageLabelWidth)
	d.containerIPAddrField.SetBackgroundColor(bgColor)
	d.containerIPAddrField.SetLabelColor(style.DialogFgColor)
	d.containerIPAddrField.SetFieldBackgroundColor(inputFieldBgColor)

	// mac field
	d.containerMacAddrField.SetLabel("mac address:")
	d.containerMacAddrField.SetLabelWidth(networkingPageLabelWidth)
	d.containerMacAddrField.SetBackgroundColor(bgColor)
	d.containerMacAddrField.SetLabelColor(style.DialogFgColor)
	d.containerMacAddrField.SetFieldBackgroundColor(inputFieldBgColor)

	// network field
	d.containerNetworkField.SetLabel("network:")
	d.containerNetworkField.SetLabelWidth(networkingPageLabelWidth)
	d.containerNetworkField.SetBackgroundColor(bgColor)
	d.containerNetworkField.SetLabelColor(style.DialogFgColor)
	d.containerNetworkField.SetListStyles(ddUnselectedStyle, ddselectedStyle)
	d.containerNetworkField.SetFieldBackgroundColor(inputFieldBgColor)

	// network settings page
	d.networkingPage.SetDirection(tview.FlexRow)
	d.networkingPage.AddItem(d.containerHostnameField, 1, 0, true)
	d.networkingPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.networkingPage.AddItem(d.containerIPAddrField, 1, 0, true)
	d.networkingPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.networkingPage.AddItem(d.containerMacAddrField, 1, 0, true)
	d.networkingPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.networkingPage.AddItem(d.containerNetworkField, 1, 0, true)
	d.networkingPage.SetBackgroundColor(bgColor)
}

func (d *ContainerCreateDialog) setupPortsPageUI() {
	bgColor := style.DialogBgColor
	inputFieldBgColor := style.InputFieldBgColor
	portPageLabelWidth := 15

	inputFieldItems := []struct {
		label  string
		widget *tview.InputField
	}{
		{label: "publish ports:", widget: d.containerPortPublishField},
		{label: "expose ports:", widget: d.containerPortExposeField},
	}

	for _, inputField := range inputFieldItems {
		inputField.widget.SetLabel(inputField.label)
		inputField.widget.SetLabelWidth(portPageLabelWidth)
		inputField.widget.SetBackgroundColor(bgColor)
		inputField.widget.SetLabelColor(style.DialogFgColor)
		inputField.widget.SetFieldBackgroundColor(inputFieldBgColor)
	}

	// publish all field
	d.ContainerPortPublishAllField.SetLabel("publish all ")
	d.ContainerPortPublishAllField.SetLabelWidth(portPageLabelWidth)
	d.ContainerPortPublishAllField.SetBackgroundColor(bgColor)
	d.ContainerPortPublishAllField.SetLabelColor(style.DialogFgColor)
	d.ContainerPortPublishAllField.SetFieldBackgroundColor(inputFieldBgColor)

	d.portPage.SetDirection(tview.FlexRow)
	d.portPage.AddItem(d.containerPortPublishField, 1, 0, true)
	d.portPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.portPage.AddItem(d.ContainerPortPublishAllField, 1, 0, true)
	d.portPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.portPage.AddItem(d.containerPortExposeField, 1, 0, true)
	d.portPage.SetBackgroundColor(bgColor)
}

func (d *ContainerCreateDialog) setupSecurityPageUI() {
	bgColor := style.DialogBgColor
	inputFieldBgColor := style.InputFieldBgColor
	securityOptsLabelWidth := 10

	// selinux label
	d.containerSecLabelField.SetLabel("label:")
	d.containerSecLabelField.SetLabelWidth(securityOptsLabelWidth)
	d.containerSecLabelField.SetBackgroundColor(bgColor)
	d.containerSecLabelField.SetLabelColor(style.DialogFgColor)
	d.containerSecLabelField.SetFieldBackgroundColor(inputFieldBgColor)

	// apparmor
	d.containerSecApparmorField.SetLabel("apparmor:")
	d.containerSecApparmorField.SetLabelWidth(securityOptsLabelWidth)
	d.containerSecApparmorField.SetBackgroundColor(bgColor)
	d.containerSecApparmorField.SetLabelColor(style.DialogFgColor)
	d.containerSecApparmorField.SetFieldBackgroundColor(inputFieldBgColor)

	// seccomp
	d.containerSeccompField.SetLabel("seccomp:")
	d.containerSeccompField.SetLabelWidth(securityOptsLabelWidth)
	d.containerSeccompField.SetBackgroundColor(bgColor)
	d.containerSeccompField.SetLabelColor(style.DialogFgColor)
	d.containerSeccompField.SetFieldBackgroundColor(inputFieldBgColor)

	// mask
	d.containerSecMaskField.SetLabel("mask:")
	d.containerSecMaskField.SetLabelWidth(securityOptsLabelWidth)
	d.containerSecMaskField.SetBackgroundColor(bgColor)
	d.containerSecMaskField.SetLabelColor(style.DialogFgColor)
	d.containerSecMaskField.SetFieldBackgroundColor(inputFieldBgColor)

	// unmask
	d.containerSecUnmaskField.SetLabel("unmask:")
	d.containerSecUnmaskField.SetLabelWidth(securityOptsLabelWidth)
	d.containerSecUnmaskField.SetBackgroundColor(bgColor)
	d.containerSecUnmaskField.SetLabelColor(style.DialogFgColor)
	d.containerSecUnmaskField.SetFieldBackgroundColor(inputFieldBgColor)

	// no-new-privileges
	d.containerSecNoNewPrivField.SetLabel("no new privileges ")
	d.containerSecNoNewPrivField.SetBackgroundColor(bgColor)
	d.containerSecNoNewPrivField.SetLabelColor(style.DialogFgColor)
	d.containerSecNoNewPrivField.SetBackgroundColor(bgColor)
	d.containerSecNoNewPrivField.SetLabelColor(style.DialogFgColor)
	d.containerSecNoNewPrivField.SetFieldBackgroundColor(inputFieldBgColor)

	// security options page
	d.securityOptsPage.SetDirection(tview.FlexRow)
	d.securityOptsPage.AddItem(d.containerSecLabelField, 1, 0, true)
	d.securityOptsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.securityOptsPage.AddItem(d.containerSecApparmorField, 1, 0, true)
	d.securityOptsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.securityOptsPage.AddItem(d.containerSeccompField, 1, 0, true)
	d.securityOptsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.securityOptsPage.AddItem(d.containerSecMaskField, 1, 0, true)
	d.securityOptsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.securityOptsPage.AddItem(d.containerSecUnmaskField, 1, 0, true)
	d.securityOptsPage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.securityOptsPage.AddItem(d.containerSecNoNewPrivField, 1, 0, true)
}

func (d *ContainerCreateDialog) setupVolumePageUI() {
	bgColor := style.DialogBgColor
	ddUnselectedStyle := style.DropDownUnselected
	ddselectedStyle := style.DropDownSelected
	inputFieldBgColor := style.InputFieldBgColor
	volumePageLabelWidth := 14

	// volume
	d.containerVolumeField.SetLabel("volume:")
	d.containerVolumeField.SetLabelWidth(volumePageLabelWidth)
	d.containerVolumeField.SetBackgroundColor(bgColor)
	d.containerVolumeField.SetLabelColor(style.DialogFgColor)
	d.containerVolumeField.SetFieldBackgroundColor(inputFieldBgColor)

	// image volume
	d.containerImageVolumeField.SetLabel("image volume:")
	d.containerImageVolumeField.SetLabelWidth(volumePageLabelWidth)
	d.containerImageVolumeField.SetBackgroundColor(bgColor)
	d.containerImageVolumeField.SetLabelColor(style.DialogFgColor)
	d.containerImageVolumeField.SetListStyles(ddUnselectedStyle, ddselectedStyle)
	d.containerImageVolumeField.SetFieldBackgroundColor(inputFieldBgColor)

	// mounts
	d.containerMountField.SetLabel("mount:")
	d.containerMountField.SetLabelWidth(volumePageLabelWidth)
	d.containerMountField.SetBackgroundColor(bgColor)
	d.containerMountField.SetLabelColor(style.DialogFgColor)
	d.containerMountField.SetFieldBackgroundColor(inputFieldBgColor)

	// volume settings page
	d.volumePage.SetDirection(tview.FlexRow)
	d.volumePage.AddItem(d.containerVolumeField, 1, 0, true)
	d.volumePage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.volumePage.AddItem(d.containerImageVolumeField, 1, 0, true)
	d.volumePage.AddItem(utils.EmptyBoxSpace(bgColor), 1, 0, true)
	d.volumePage.AddItem(d.containerMountField, 1, 0, true)
	d.volumePage.SetBackgroundColor(bgColor)
}

// Display displays this primitive.
func (d *ContainerCreateDialog) Display() {
	d.display = true
	d.initData()
	d.focusElement = createCategoryPagesFocus
}

// IsDisplay returns true if primitive is shown.
func (d *ContainerCreateDialog) IsDisplay() bool {
	return d.display
}

// Hide stops displaying this primitive.
func (d *ContainerCreateDialog) Hide() {
	d.display = false
}

// HasFocus returns whether or not this primitive has focus.
func (d *ContainerCreateDialog) HasFocus() bool {
	if d.categories.HasFocus() || d.categoryPages.HasFocus() {
		return true
	}

	return d.Box.HasFocus() || d.form.HasFocus()
}

// dropdownHasFocus returns true if container create dialog dropdown primitives
// has focus.
func (d *ContainerCreateDialog) dropdownHasFocus() bool {
	if d.containerImageField.HasFocus() || d.containerPodField.HasFocus() {
		return true
	}

	if d.containerNetworkField.HasFocus() || d.containerImageVolumeField.HasFocus() {
		return true
	}

	return d.containerHealthOnFailureField.HasFocus()
}

// Focus is called when this primitive receives focus.
func (d *ContainerCreateDialog) Focus(delegate func(p tview.Primitive)) { //nolint:gocyclo,cyclop,maintidx
	switch d.focusElement {
	// form has focus
	case createContainerFormFocus:
		button := d.form.GetButton(d.form.GetButtonCount() - 1)
		button.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				d.focusElement = createCategoriesFocus // category text view

				d.Focus(delegate)
				d.form.SetFocus(0)

				return nil
			}

			if event.Key() == tcell.KeyEnter {
				return nil
			}

			return event
		})

		delegate(d.form)
	// category text view
	case createCategoriesFocus:
		d.categories.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyTab {
				d.focusElement = createCategoryPagesFocus // category page view
				d.Focus(delegate)

				return nil
			}

			// scroll between categories
			event = utils.ParseKeyEventKey(event)
			if event.Key() == tcell.KeyDown {
				d.nextCategory()
			}

			if event.Key() == tcell.KeyUp {
				d.previousCategory()
			}

			return event
		})

		delegate(d.categories)
	// container info page
	case createContainerNameFieldFocus:
		delegate(d.containerNameField)
	case createContainerImageFieldFocus:
		delegate(d.containerImageField)
	case createcontainerPodFieldFocis:
		delegate(d.containerPodField)
	case createContainerLabelsFieldFocus:
		delegate(d.containerLabelsField)
	case createContainerRemoveFieldFocus:
		delegate(d.containerRemoveField)
	case createContainerPrivilegedFieldFocus:
		delegate(d.containerPrivilegedField)
	case createContainerTimeoutFieldFocus:
		delegate(d.containerTimeoutField)
	// environment options page
	case createContainerWorkDirFieldFocus:
		delegate(d.containerWorkDirField)
	case createContainerEnvVarsFieldFocus:
		delegate(d.containerEnvVarsField)
	case createContainerEnvFileFieldFocus:
		delegate(d.containerEnvFileField)
	case createContainerEnvMergeFieldFocus:
		delegate(d.containerEnvMergeField)
	case createContainerUnsetEnvFieldFocus:
		delegate(d.containerUnsetEnvField)
	case createContainerEnvHostFieldFocus:
		delegate(d.containerEnvHostField)
	case createContainerUnsetEnvAllFieldFocus:
		delegate(d.containerUnsetEnvAllField)
	case createContainerUmaskFieldFocus:
		delegate(d.containerUmaskField)
	// user and groups page
	case createContainerUserFieldFocus:
		delegate(d.containerUserField)
	case createContainerHostUsersFieldFocus:
		delegate(d.containerHostUsersField)
	case createContainerPasswdEntryFieldFocus:
		delegate(d.containerPasswdEntryField)
	case createContainerGroupEntryFieldFocus:
		delegate(d.containerGroupEntryField)
	// security options page
	case createcontainerSecLabelFieldFocus:
		delegate(d.containerSecLabelField)
	case createContainerApprarmorFieldFocus:
		delegate(d.containerSecApparmorField)
	case createContainerSeccompFeildFocus:
		delegate(d.containerSeccompField)
	case createcontainerSecMaskFieldFocus:
		delegate(d.containerSecMaskField)
	case createcontainerSecUnmaskFieldFocus:
		delegate(d.containerSecUnmaskField)
	case createcontainerSecNoNewPrivFieldFocus:
		delegate(d.containerSecNoNewPrivField)
	// networking page
	case createContainerHostnameFieldFocus:
		delegate(d.containerHostnameField)
	case createContainerIPAddrFieldFocus:
		delegate(d.containerIPAddrField)
	case createContainerMacAddrFieldFocus:
		delegate(d.containerMacAddrField)
	case createContainerNetworkFieldFocus:
		delegate(d.containerNetworkField)
	// port page
	// networking page
	case createContainerPortPublishFieldFocus:
		delegate(d.containerPortPublishField)
	case createContainerPortPublishAllFieldFocus:
		delegate(d.ContainerPortPublishAllField)
	case createContainerPortExposeFieldFocus:
		delegate(d.containerPortExposeField)
	// dns page
	case createContainerDNSServersFieldFocus:
		delegate(d.containerDNSServersField)
	case createContainerDNSOptionsFieldFocus:
		delegate(d.containerDNSOptionsField)
	case createContainerDNSSearchFieldFocus:
		delegate(d.containerDNSSearchField)
	// volume page
	case createContainerVolumeFieldFocus:
		delegate(d.containerVolumeField)
	case createContainerImageVolumeFieldFocus:
		delegate(d.containerImageVolumeField)
	case createContainerMountFieldFocus:
		delegate(d.containerMountField)
	// health page
	case containerHealthCmdFieldFocus:
		delegate(d.containerHealthCmdField)
	case containerHealthStartupCmdFieldFocus:
		delegate(d.containerHealthStartupCmdField)
	case containerHealthOnFailureFieldFocus:
		delegate(d.containerHealthOnFailureField)
	case containerHealthIntervalFieldFocus:
		delegate(d.containerHealthIntervalField)
	case containerHealthStartupIntervalFieldFocus:
		delegate(d.containerHealthStartupIntervalField)
	case containerHealthTimeoutFieldFocus:
		delegate(d.containerHealthTimeoutField)
	case containerHealthStartupTimeoutFieldFocus:
		delegate(d.containerHealthStartupTimeoutField)
	case containerHealthRetriesFieldFocus:
		delegate(d.containerHealthRetriesField)
	case containerHealthStartupRetriesFieldFocus:
		delegate(d.containerHealthStartupRetriesField)
	case containerHealthStartPeriodFieldFocus:
		delegate(d.containerHealthStartPeriodField)
	case containerHealthStartupSuccessFieldFocus:
		delegate(d.containerHealthStartupSuccessField)
	// category page
	case createCategoryPagesFocus:
		delegate(d.categoryPages)
	}
}

func (d *ContainerCreateDialog) initCustomInputHanlers() {
	// pod name dropdown
	d.containerPodField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = utils.ParseKeyEventKey(event)

		return event
	})

	// container image volume
	d.containerImageField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = utils.ParseKeyEventKey(event)

		return event
	})

	// container network dropdown
	d.containerNetworkField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = utils.ParseKeyEventKey(event)

		return event
	})

	// container image volume dropdown
	d.containerImageVolumeField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		event = utils.ParseKeyEventKey(event)

		return event
	})
}

// InputHandler returns input handler function for this primitive.
func (d *ContainerCreateDialog) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) { //nolint:gocognit,lll,cyclop,gocyclo
	return d.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		log.Debug().Msgf("container create dialog: event %v received", event)

		if event.Key() == tcell.KeyEsc && !d.dropdownHasFocus() {
			d.cancelHandler()

			return
		}

		if d.containerInfoPage.HasFocus() {
			if handler := d.containerInfoPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setContainerInfoPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.environmentPage.HasFocus() {
			if handler := d.environmentPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setEnvironmentPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.userGroupsPage.HasFocus() {
			if handler := d.userGroupsPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setUserGroupsPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.dnsPage.HasFocus() {
			if handler := d.dnsPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setDNSSettingsPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.networkingPage.HasFocus() {
			if handler := d.networkingPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setNetworkSettingsPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.portPage.HasFocus() {
			if handler := d.portPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setPortPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.securityOptsPage.HasFocus() {
			if handler := d.securityOptsPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setSecurityOptionsPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.volumePage.HasFocus() {
			if handler := d.volumePage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setVolumeSettingsPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.healthPage.HasFocus() {
			if handler := d.healthPage.InputHandler(); handler != nil {
				if event.Key() == tcell.KeyTab {
					d.setHealthSettingsPageNextFocus()
				}

				handler(event, setFocus)

				return
			}
		}

		if d.categories.HasFocus() {
			if categroryHandler := d.categories.InputHandler(); categroryHandler != nil {
				categroryHandler(event, setFocus)

				return
			}
		}

		if d.form.HasFocus() { //nolint:nestif
			if formHandler := d.form.InputHandler(); formHandler != nil {
				if event.Key() == tcell.KeyEnter {
					enterButton := d.form.GetButton(d.form.GetButtonCount() - 1)
					if enterButton.HasFocus() {
						d.createHandler()
					}
				}

				formHandler(event, setFocus)

				return
			}
		}
	})
}

// SetRect set rects for this primitive.
func (d *ContainerCreateDialog) SetRect(x, y, width, height int) {
	if width > containerCreateDialogMaxWidth {
		emptySpace := (width - containerCreateDialogMaxWidth) / 2 //nolint:gomnd
		x += emptySpace
		width = containerCreateDialogMaxWidth
	}

	if height > containerCreateDialogHeight {
		emptySpace := (height - containerCreateDialogHeight) / 2 //nolint:gomnd
		y += emptySpace
		height = containerCreateDialogHeight
	}

	d.Box.SetRect(x, y, width, height)
}

// Draw draws this primitive onto the screen.
func (d *ContainerCreateDialog) Draw(screen tcell.Screen) {
	if !d.display {
		return
	}

	d.Box.DrawForSubclass(screen, d)

	x, y, width, height := d.Box.GetInnerRect()

	d.layout.SetRect(x, y, width, height)
	d.layout.Draw(screen)
}

// SetCancelFunc sets form cancel button selected function.
func (d *ContainerCreateDialog) SetCancelFunc(handler func()) *ContainerCreateDialog {
	d.cancelHandler = handler
	cancelButton := d.form.GetButton(d.form.GetButtonCount() - 2) //nolint:gomnd

	cancelButton.SetSelectedFunc(handler)

	return d
}

// SetCreateFunc sets form create button selected function.
func (d *ContainerCreateDialog) SetCreateFunc(handler func()) *ContainerCreateDialog {
	d.createHandler = handler
	enterButton := d.form.GetButton(d.form.GetButtonCount() - 1)

	enterButton.SetSelectedFunc(handler)

	return d
}

func (d *ContainerCreateDialog) setActiveCategory(index int) {
	fgColor := style.DialogFgColor
	bgColor := style.ButtonBgColor
	ctgTextColor := style.GetColorHex(fgColor)
	ctgBgColor := style.GetColorHex(bgColor)

	d.activePageIndex = index

	d.categories.Clear()

	var ctgList []string

	alignedList, _ := utils.AlignStringListWidth(d.categoryLabels)

	for i := 0; i < len(alignedList); i++ {
		if i == index {
			ctgList = append(ctgList, fmt.Sprintf("[%s:%s:b]-> %s ", ctgTextColor, ctgBgColor, alignedList[i]))

			continue
		}

		ctgList = append(ctgList, fmt.Sprintf("[-:-:-]   %s ", alignedList[i]))
	}

	d.categories.SetText(strings.Join(ctgList, "\n"))

	// switch the page
	d.categoryPages.SwitchToPage(d.categoryLabels[index])
}

func (d *ContainerCreateDialog) nextCategory() {
	activePage := d.activePageIndex
	if d.activePageIndex < len(d.categoryLabels)-1 {
		activePage++

		d.setActiveCategory(activePage)

		return
	}

	d.setActiveCategory(0)
}

func (d *ContainerCreateDialog) previousCategory() {
	activePage := d.activePageIndex
	if d.activePageIndex > 0 {
		activePage--

		d.setActiveCategory(activePage)

		return
	}

	d.setActiveCategory(len(d.categoryLabels) - 1)
}

func (d *ContainerCreateDialog) initData() {
	// get available images
	imgList, _ := images.List()
	d.imageList = imgList
	imgOptions := []string{""}

	for i := 0; i < len(d.imageList); i++ {
		if d.imageList[i].ID == "<none>" {
			imgOptions = append(imgOptions, d.imageList[i].ID)

			continue
		}

		imgname := d.imageList[i].Repository + ":" + d.imageList[i].Tag
		imgOptions = append(imgOptions, imgname)
	}

	// get available pods
	podOptions := []string{""}
	podList, _ := pods.List()
	d.podList = podList

	for i := 0; i < len(podList); i++ {
		podOptions = append(podOptions, podList[i].Name)
	}

	// get available networks
	networkOptions := []string{""}
	networkList, _ := networks.List()

	for i := 0; i < len(networkList); i++ {
		networkOptions = append(networkOptions, networkList[i][1])
	}

	// get available volumes
	imageVolumeOptions := []string{"", "ignore", "tmpfs", "anonymous"}
	volumeOptions := []string{""}
	volList, _ := volumes.List()

	for i := 0; i < len(volList); i++ {
		volumeOptions = append(volumeOptions, volList[i].Name)
	}

	d.setActiveCategory(0)
	// container category
	d.containerNameField.SetText("")
	d.containerImageField.SetOptions(imgOptions, nil)
	d.containerImageField.SetCurrentOption(0)
	d.containerPodField.SetOptions(podOptions, nil)
	d.containerPodField.SetCurrentOption(0)
	d.containerLabelsField.SetText("")
	d.containerRemoveField.SetChecked(false)
	d.containerPrivilegedField.SetChecked(false)
	d.containerTimeoutField.SetText("")

	// environment category
	d.containerWorkDirField.SetText("")
	d.containerEnvVarsField.SetText("")
	d.containerEnvFileField.SetText("")
	d.containerEnvMergeField.SetText("")
	d.containerUnsetEnvField.SetText("")
	d.containerEnvHostField.SetChecked(false)
	d.containerUnsetEnvAllField.SetChecked(false)
	d.containerUmaskField.SetText("")

	// user and groups category
	d.containerUserField.SetText("")
	d.containerHostUsersField.SetText("")
	d.containerPasswdEntryField.SetText("")
	d.containerGroupEntryField.SetText("")

	// dns settings category
	d.containerDNSServersField.SetText("")
	d.containerDNSSearchField.SetText("")
	d.containerDNSOptionsField.SetText("")

	// health options category
	d.containerHealthCmdField.SetText("")
	d.containerHealthStartupCmdField.SetText("")
	d.containerHealthOnFailureField.SetCurrentOption(0)
	d.containerHealthIntervalField.SetText("")
	d.containerHealthStartupIntervalField.SetText("")
	d.containerHealthTimeoutField.SetText("")
	d.containerHealthStartupTimeoutField.SetText("")
	d.containerHealthRetriesField.SetText("")
	d.containerHealthStartupRetriesField.SetText("")
	d.containerHealthStartPeriodField.SetText("")
	d.containerHealthStartupSuccessField.SetText("")

	// network settings category
	d.containerHostnameField.SetText("")
	d.containerIPAddrField.SetText("")
	d.containerMacAddrField.SetText("")
	d.containerNetworkField.SetOptions(networkOptions, nil)
	d.containerNetworkField.SetCurrentOption(0)

	// ports settings category
	d.containerPortPublishField.SetText("")
	d.ContainerPortPublishAllField.SetChecked(false)
	d.containerPortExposeField.SetText("")

	// security options category
	d.containerSecLabelField.SetText("")
	d.containerSecApparmorField.SetText("")
	d.containerSeccompField.SetText("")
	d.containerSecMaskField.SetText("")
	d.containerSecUnmaskField.SetText("")
	d.containerSecNoNewPrivField.SetChecked(false)

	// volumes options category
	d.containerVolumeField.SetText("")
	d.containerMountField.SetText("")
	d.containerImageVolumeField.SetOptions(imageVolumeOptions, nil)
	d.containerImageVolumeField.SetCurrentOption(0)
}

func (d *ContainerCreateDialog) setPortPageNextFocus() {
	if d.containerPortPublishField.HasFocus() {
		d.focusElement = createContainerPortPublishAllFieldFocus

		return
	}

	if d.ContainerPortPublishAllField.HasFocus() {
		d.focusElement = createContainerPortExposeFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setContainerInfoPageNextFocus() {
	if d.containerNameField.HasFocus() {
		d.focusElement = createContainerImageFieldFocus

		return
	}

	if d.containerImageField.HasFocus() {
		d.focusElement = createcontainerPodFieldFocis

		return
	}

	if d.containerPodField.HasFocus() {
		d.focusElement = createContainerLabelsFieldFocus

		return
	}

	if d.containerLabelsField.HasFocus() {
		d.focusElement = createContainerPrivilegedFieldFocus

		return
	}

	if d.containerPrivilegedField.HasFocus() {
		d.focusElement = createContainerRemoveFieldFocus

		return
	}

	if d.containerRemoveField.HasFocus() {
		d.focusElement = createContainerTimeoutFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setEnvironmentPageNextFocus() {
	if d.containerWorkDirField.HasFocus() {
		d.focusElement = createContainerEnvVarsFieldFocus

		return
	}

	if d.containerEnvVarsField.HasFocus() {
		d.focusElement = createContainerEnvFileFieldFocus

		return
	}

	if d.containerEnvFileField.HasFocus() {
		d.focusElement = createContainerEnvMergeFieldFocus

		return
	}

	if d.containerEnvMergeField.HasFocus() {
		d.focusElement = createContainerUnsetEnvFieldFocus

		return
	}

	if d.containerUnsetEnvField.HasFocus() {
		d.focusElement = createContainerEnvHostFieldFocus

		return
	}

	if d.containerEnvHostField.HasFocus() {
		d.focusElement = createContainerUnsetEnvAllFieldFocus

		return
	}

	if d.containerUnsetEnvAllField.HasFocus() {
		d.focusElement = createContainerUmaskFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setUserGroupsPageNextFocus() {
	if d.containerUserField.HasFocus() {
		d.focusElement = createContainerHostUsersFieldFocus

		return
	}

	if d.containerHostUsersField.HasFocus() {
		d.focusElement = createContainerPasswdEntryFieldFocus

		return
	}

	if d.containerPasswdEntryField.HasFocus() {
		d.focusElement = createContainerGroupEntryFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setSecurityOptionsPageNextFocus() {
	if d.containerSecLabelField.HasFocus() {
		d.focusElement = createContainerApprarmorFieldFocus

		return
	}

	if d.containerSecApparmorField.HasFocus() {
		d.focusElement = createContainerSeccompFeildFocus

		return
	}

	if d.containerSeccompField.HasFocus() {
		d.focusElement = createcontainerSecMaskFieldFocus

		return
	}

	if d.containerSecMaskField.HasFocus() {
		d.focusElement = createcontainerSecUnmaskFieldFocus

		return
	}

	if d.containerSecUnmaskField.HasFocus() {
		d.focusElement = createcontainerSecNoNewPrivFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setNetworkSettingsPageNextFocus() {
	if d.containerHostnameField.HasFocus() {
		d.focusElement = createContainerIPAddrFieldFocus

		return
	}

	if d.containerIPAddrField.HasFocus() {
		d.focusElement = createContainerMacAddrFieldFocus

		return
	}

	if d.containerMacAddrField.HasFocus() {
		d.focusElement = createContainerNetworkFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setDNSSettingsPageNextFocus() {
	if d.containerDNSServersField.HasFocus() {
		d.focusElement = createContainerDNSOptionsFieldFocus

		return
	}

	if d.containerDNSOptionsField.HasFocus() {
		d.focusElement = createContainerDNSSearchFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setHealthSettingsPageNextFocus() { //nolint:cyclop
	if d.containerHealthCmdField.HasFocus() {
		d.focusElement = containerHealthStartupCmdFieldFocus

		return
	}

	if d.containerHealthStartupCmdField.HasFocus() {
		d.focusElement = containerHealthOnFailureFieldFocus

		return
	}

	if d.containerHealthOnFailureField.HasFocus() {
		d.focusElement = containerHealthStartupSuccessFieldFocus

		return
	}

	if d.containerHealthStartupSuccessField.HasFocus() {
		d.focusElement = containerHealthStartPeriodFieldFocus

		return
	}

	if d.containerHealthStartPeriodField.HasFocus() {
		d.focusElement = containerHealthIntervalFieldFocus

		return
	}

	if d.containerHealthIntervalField.HasFocus() {
		d.focusElement = containerHealthStartupIntervalFieldFocus

		return
	}

	if d.containerHealthStartupIntervalField.HasFocus() {
		d.focusElement = containerHealthRetriesFieldFocus

		return
	}

	if d.containerHealthRetriesField.HasFocus() {
		d.focusElement = containerHealthStartupRetriesFieldFocus

		return
	}

	if d.containerHealthStartupRetriesField.HasFocus() {
		d.focusElement = containerHealthTimeoutFieldFocus

		return
	}

	if d.containerHealthTimeoutField.HasFocus() {
		d.focusElement = containerHealthStartupTimeoutFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

func (d *ContainerCreateDialog) setVolumeSettingsPageNextFocus() {
	if d.containerVolumeField.HasFocus() {
		d.focusElement = createContainerImageVolumeFieldFocus

		return
	}

	if d.containerImageVolumeField.HasFocus() {
		d.focusElement = createContainerMountFieldFocus

		return
	}

	d.focusElement = createContainerFormFocus
}

// ContainerCreateOptions returns new network options.
func (d *ContainerCreateDialog) ContainerCreateOptions() containers.CreateOptions { //nolint:cyclop,gocognit
	var (
		labels           []string
		imageID          string
		podID            string
		dnsServers       []string
		dnsOptions       []string
		dnsSearchDomains []string
		publish          []string
		expose           []string
		imageVolume      string
		selinuxOpts      []string
		envVars          []string
		envFile          []string
		envMerge         []string
		unsetEnv         []string
		hostUsers        []string
	)

	for _, label := range strings.Split(d.containerLabelsField.GetText(), " ") {
		if label != "" {
			labels = append(labels, label)
		}
	}

	selectedImageIndex, _ := d.containerImageField.GetCurrentOption()
	if len(d.imageList) > 0 && selectedImageIndex > 0 {
		imageID = d.imageList[selectedImageIndex-1].ID
	}

	selectedPodIndex, _ := d.containerPodField.GetCurrentOption()
	if len(d.podList) > 0 && selectedPodIndex > 0 {
		podID = d.podList[selectedPodIndex-1].Id
	}

	// ports
	for _, p := range strings.Split(d.containerPortPublishField.GetText(), " ") {
		if p != "" {
			publish = append(publish, p)
		}
	}

	for _, e := range strings.Split(d.containerPortExposeField.GetText(), " ") {
		if e != "" {
			expose = append(expose, e)
		}
	}

	// DNS setting
	for _, dns := range strings.Split(d.containerDNSServersField.GetText(), " ") {
		if dns != "" {
			dnsServers = append(dnsServers, dns)
		}
	}

	for _, do := range strings.Split(d.containerDNSOptionsField.GetText(), " ") {
		if do != "" {
			dnsOptions = append(dnsOptions, do)
		}
	}

	for _, ds := range strings.Split(d.containerDNSSearchField.GetText(), " ") {
		if ds != "" {
			dnsSearchDomains = append(dnsSearchDomains, ds)
		}
	}

	_, imageVolume = d.containerImageVolumeField.GetCurrentOption()

	// security options
	for _, selinuxLabel := range strings.Split(d.containerSecLabelField.GetText(), " ") {
		if selinuxLabel != "" {
			selinuxOpts = append(selinuxOpts, selinuxLabel)
		}
	}

	// health check
	_, healthOnFailure := d.containerHealthOnFailureField.GetCurrentOption()

	// env vars
	for _, evar := range strings.Split(d.containerEnvVarsField.GetText(), " ") {
		if evar != "" {
			envVars = append(envVars, evar)
		}
	}

	// env file
	for _, efile := range strings.Split(d.containerEnvFileField.GetText(), " ") {
		if efile != "" {
			envFile = append(envFile, efile)
		}
	}

	// env merge
	for _, emerge := range strings.Split(d.containerEnvMergeField.GetText(), " ") {
		if emerge != "" {
			envMerge = append(envMerge, emerge)
		}
	}

	// unset env
	for _, eunset := range strings.Split(d.containerUnsetEnvField.GetText(), " ") {
		if eunset != "" {
			unsetEnv = append(unsetEnv, eunset)
		}
	}

	// host users
	for _, huser := range strings.Split(d.containerHostUsersField.GetText(), " ") {
		if huser != "" {
			hostUsers = append(hostUsers, huser)
		}
	}

	_, network := d.containerNetworkField.GetCurrentOption()
	opts := containers.CreateOptions{
		Name:                  d.containerNameField.GetText(),
		Image:                 imageID,
		Pod:                   podID,
		Labels:                labels,
		Remove:                d.containerRemoveField.IsChecked(),
		Privileged:            d.containerPrivilegedField.IsChecked(),
		Timeout:               d.containerTimeoutField.GetText(),
		WorkDir:               d.containerWorkDirField.GetText(),
		EnvVars:               envVars,
		EnvFile:               envFile,
		EnvMerge:              envMerge,
		UnsetEnv:              unsetEnv,
		EnvHost:               d.containerEnvHostField.IsChecked(),
		UnsetEnvAll:           d.containerUnsetEnvAllField.IsChecked(),
		Umask:                 d.containerUmaskField.GetText(),
		User:                  d.containerUserField.GetText(),
		HostUsers:             hostUsers,
		PasswdEntry:           d.containerPasswdEntryField.GetText(),
		GroupEntry:            d.containerGroupEntryField.GetText(),
		Hostname:              d.containerHostnameField.GetText(),
		MacAddress:            d.containerMacAddrField.GetText(),
		IPAddress:             d.containerIPAddrField.GetText(),
		Network:               network,
		Publish:               publish,
		Expose:                expose,
		PublishAll:            d.ContainerPortPublishAllField.IsChecked(),
		DNSServer:             dnsServers,
		DNSOptions:            dnsOptions,
		DNSSearchDomain:       dnsSearchDomains,
		Volume:                d.containerVolumeField.GetText(),
		ImageVolume:           imageVolume,
		Mount:                 d.containerMountField.GetText(),
		SelinuxOpts:           selinuxOpts,
		ApparmorProfile:       d.containerSecApparmorField.GetText(),
		Seccomp:               d.containerSeccompField.GetText(),
		SecNoNewPriv:          d.containerSecNoNewPrivField.IsChecked(),
		SecMask:               d.containerSecMaskField.GetText(),
		SecUnmask:             d.containerSecUnmaskField.GetText(),
		HealthCmd:             strings.TrimSpace(d.containerHealthCmdField.GetText()),
		HealthInterval:        strings.TrimSpace(d.containerHealthIntervalField.GetText()),
		HealthRetries:         strings.TrimSpace(d.containerHealthRetriesField.GetText()),
		HealthStartPeroid:     strings.TrimSpace(d.containerHealthStartPeriodField.GetText()),
		HealthTimeout:         strings.TrimSpace(d.containerHealthTimeoutField.GetText()),
		HealthOnFailure:       strings.TrimSpace(healthOnFailure),
		HealthStartupCmd:      strings.TrimSpace(d.containerHealthStartupCmdField.GetText()),
		HealthStartupInterval: strings.TrimSpace(d.containerHealthStartupIntervalField.GetText()),
		HealthStartupRetries:  strings.TrimSpace(d.containerHealthStartupRetriesField.GetText()),
		HealthStartupSuccess:  strings.TrimSpace(d.containerHealthStartupSuccessField.GetText()),
		HealthStartupTimeout:  strings.TrimSpace(d.containerHealthStartupTimeoutField.GetText()),
	}

	return opts
}
