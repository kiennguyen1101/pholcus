package gui

import (
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/henrylee2cn/pholcus/app"
	"github.com/henrylee2cn/pholcus/config"
	"github.com/henrylee2cn/pholcus/logs"
)

var serverCount int

func serverWindow() {
	mw.Close()

	if err := (MainWindow{
		AssignTo: &mw,
		DataBinder: DataBinder{
			AssignTo:       &db,
			DataSource:     Input,
			ErrorPresenter: ErrorPresenterRef{&ep},
		},
		Title:   config.FULL_NAME + "                                                          【 Operating mode -> Server 】",
		MinSize: Size{1100, 700},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{

			Composite{
				AssignTo: &setting,
				Layout:   Grid{Columns: 2},
				Children: []Widget{
					// 任务列表
					TableView{
						ColumnSpan:            1,
						MinSize:               Size{550, 450},
						AlternatingRowBGColor: walk.RGB(255, 255, 224),
						CheckBoxes:            true,
						ColumnsOrderable:      true,
						Columns: []TableViewColumn{
							{Title: "#", Width: 45},
							{Title: "Task", Width: 110 /*, Format: "%.2f", Alignment: AlignFar*/},
							{Title: "Description", Width: 370},
						},
						Model: spiderMenu,
					},

					VSplitter{
						ColumnSpan: 1,
						MinSize:    Size{550, 450},
						Children: []Widget{

							VSplitter{
								Children: []Widget{
									Label{
										Text: "Custom configuration（Multi-task, please pack a layer of "<>"）",
									},
									LineEdit{
										Text: Bind("Keyins"),
									},
								},
							},

							VSplitter{
								Children: []Widget{
									Label{
										Text: "*Acquisition limit (default limit URL number)：",
									},
									NumberEdit{
										Value:    Bind("Limit"),
										Suffix:   "",
										Decimals: 0,
									},
								},
							},

							VSplitter{
								Children: []Widget{
									Label{
										Text: "*Concurrent coroutine：（1~99999）",
									},
									NumberEdit{
										Value:    Bind("ThreadNum", Range{1, 99999}),
										Suffix:   "",
										Decimals: 0,
									},
								},
							},

							VSplitter{
								Children: []Widget{
									Label{
										Text: "*Batch output size: (1~5,000,000 data)",
									},
									NumberEdit{
										Value:    Bind("DockerCap", Range{1, 5000000}),
										Suffix:   "",
										Decimals: 0,
									},
								},
							},

							VSplitter{
								Children: []Widget{
									Label{
										Text: "*Pause duration:",
									},
									ComboBox{
										Value:         Bind("Pausetime", SelRequired{}),
										DisplayMember: "Key",
										BindingMember: "Int64",
										Model:         GuiOpt.Pausetime,
									},
								},
							},

							VSplitter{
								Children: []Widget{
									Label{
										Text: "*Proxy IP replacement frequency:",
									},
									ComboBox{
										Value:         Bind("ProxyMinute", SelRequired{}),
										DisplayMember: "Key",
										BindingMember: "Int64",
										Model:         GuiOpt.ProxyMinute,
									},
								},
							},

							RadioButtonGroupBox{
								ColumnSpan: 1,
								Title:      "*Output Method",
								Layout:     HBox{},
								DataMember: "OutType",
								Buttons:    outputList,
							},
						},
					},
				},
			},

			Composite{
				Layout: HBox{},
				Children: []Widget{

					// 必填项错误检查
					LineErrorPresenter{
						AssignTo: &ep,
					},

					HSplitter{
						MaxSize: Size{220, 50},
						Children: []Widget{
							Label{
								Text: "Inherit and save a successful record",
							},
							CheckBox{
								Checked: Bind("SuccessInherit"),
							},
						},
					},

					HSplitter{
						MaxSize: Size{220, 50},
						Children: []Widget{
							Label{
								Text: "Inherit and save the failed record",
							},
							CheckBox{
								Checked: Bind("FailureInherit"),
							},
						},
					},

					PushButton{
						MinSize:   Size{90, 0},
						Text:      serverBtnTxt(),
						AssignTo:  &runStopBtn,
						OnClicked: serverStart,
					},
				},
			},
		},
	}.Create()); err != nil {
		panic(err)
	}

	setWindow()

	// 初始化应用
	Init()

	// 运行窗体程序
	mw.Run()
}

// 点击开始事件
func serverStart() {
	if err := db.Submit(); err != nil {
		logs.Log.Error("%v", err)
		return
	}

	// 读取任务
	Input.Spiders = spiderMenu.GetChecked()

	if len(Input.Spiders) == 0 {
		logs.Log.Warning(" *     —— Pro, the task list can't be empty~")
		return
	}

	// 记录配置信息
	SetTaskConf()

	runStopBtn.SetEnabled(false)
	runStopBtn.SetText("分发任务 (···)")

	// 重置spiders队列
	SpiderPrepare()

	// 生成分发任务
	app.LogicApp.Run()

	serverCount++

	runStopBtn.SetText(serverBtnTxt())
	runStopBtn.SetEnabled(true)
}

// 更新按钮文字
func serverBtnTxt() string {
	return "分发任务 (" + strconv.Itoa(serverCount) + ")"
}
