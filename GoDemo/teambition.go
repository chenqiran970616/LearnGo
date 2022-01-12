package main

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = ""
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "track"
)

type Result struct {
	AccomplishDate time.Time     `json:"accomplishDate"`
	AncestorIds    []interface{} `json:"ancestorIds"`
	Content        string        `json:"content"`
	Created        time.Time     `json:"created"`
	CreatorId      string        `json:"creatorId"`
	Customfields   []struct {
		CfId  string `json:"cfId"`
		Type  string `json:"type"`
		Value []struct {
			Id    string `json:"id"`
			Title string `json:"title"`
		} `json:"value"`
	} `json:"customfields"`
	DueDate      time.Time     `json:"dueDate"`
	ExecutorId   string        `json:"executorId"`
	IsDone       int           `json:"isDone"`
	Note         string        `json:"note"`
	OrgId        interface{}   `json:"orgId"`
	Participants []string      `json:"participants"`
	Priority     int           `json:"priority"`
	ProjectId    string        `json:"projectId"`
	SprintId     string        `json:"sprintId"`
	StartDate    time.Time     `json:"startDate"`
	StatusId     string        `json:"statusId"`
	TagIds       []interface{} `json:"tagIds"`
	TaskId       string        `json:"taskId"`
	TaskgroupId  string        `json:"taskgroupId"`
	TasklistId   string        `json:"tasklistId"`
	TemplateId   string        `json:"templateId"`
	UniqueId     int           `json:"uniqueId"`
	Updated      time.Time     `json:"updated"`
	Visible      string        `json:"visible"`
}

type Track struct {
	Code          int      `json:"code"`
	ErrorMessage  string   `json:"errorMessage"`
	Count         int      `json:"count"`
	NextPageToken string   `json:"nextPageToken"`
	Result        []Result `json:"result"`
}

type User struct {
	Code          int    `json:"code"`
	ErrorMessage  string `json:"errorMessage"`
	Count         int    `json:"count"`
	NextPageToken string `json:"nextPageToken"`
	Result        []struct {
		Birthday   interface{} `json:"birthday"`
		City       string      `json:"city"`
		Country    string      `json:"country"`
		Email      string      `json:"email"`
		EntryTime  interface{} `json:"entryTime"`
		IsDisabled int         `json:"isDisabled"`
		MemberId   string      `json:"memberId"`
		Name       string      `json:"name"`
		Phone      string      `json:"phone"`
		Pinyin     string      `json:"pinyin"`
		Province   string      `json:"province"`
		Py         string      `json:"py"`
		Role       int         `json:"role"`
		StaffType  string      `json:"staffType"`
		Title      string      `json:"title"`
		UserId     string      `json:"userId"`
	} `json:"result"`
}

type UserBind struct {
	Username sql.NullString `db:"username"`
	Userid   sql.NullString `db:"userid"`
}

type UserProject struct {
	Code          int    `json:"code"`
	ErrorMessage  string `json:"errorMessage"`
	Count         int    `json:"count"`
	NextPageToken string `json:"nextPageToken"`
	Result        []struct {
		AvatarUrl string    `json:"avatarUrl"`
		Email     string    `json:"email"`
		MemberId  string    `json:"memberId"`
		Name      string    `json:"name"`
		Phone     string    `json:"phone"`
		Role      int       `json:"role"`
		Updated   time.Time `json:"updated"`
		UserId    string    `json:"userId"`
	} `json:"result"`
}

func dbInit() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return nil
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
	return DB
}

func insertIssueData(DB *sql.DB, trackTitle string, creator string, trackContent string, createTime time.Time, updateTime time.Time, index int, executor string, bugLevel int, bugStatus int, project string, bugType string, accomplishDate time.Time, dueDate time.Time, bugDelay int, iteration string) {
	//uuid, err := uuid.NewV4()
	//uuidstr := uuid.String()
	//result, err := DB.Exec("insert INTO issues(id,num,title,description,status,creator,platform,project_id,create_time,custom_fields,reporter) values(?,?,?,?,?,?,?,?,?,?,?)",uuidstr,100000+num,trackTitle,trackContent,"new",creator,"Teambition","36399bc8-b0dd-4d54-82e2-f77ee259a530",createTime.Unix()*1000,"[{\"id\":\"07343f98-997b-43f8-a8f7-807364b544ba\",\"name\":\"严重程度\",\"value\":\"P0\",\"type\":\"select\",\"customData\":null},{\"id\":\"172ed88b-849e-4af6-afd9-1ccc89d0144d\",\"name\":\"创建时间\",\"value\":\"2021-08-31T16:00:00.000Z\",\"type\":\"data\",\"customData\":null},{\"id\":\"340cd4cf-fb15-4d75-a5a7-4f8c640b46c8\",\"name\":\"创建人\",\"value\":\"+"+creator+"\",\"type\":\"member\",\"customData\":null},{\"id\":\"ba9663a7-9942-4480-bc16-bafe130dc85a\",\"name\":\"bug所属\",\"value\":[\""+executor+"\"],\"type\":\"multipleMember\",\"customData\":null},{\"id\":\"d40ac7a2-e2e9-4993-ae49-75d08682f050\",\"name\":\"状态\",\"value\":\"new\",\"type\":\"select\",\"customData\":null},{\"id\":\"e5f80716-0e56-4a5f-a28b-8683a41fdbbf\",\"name\":\"处理人\",\"value\":\""+executor+"\",\"type\":\"member\",\"customData\":null}]",executor)
	result, err := DB.Exec("insert INTO issues_statistics(issueTitle,issueContent,bugOwner,creator,executor,createTime,updateTime,bugLevel,bugStatus,project,bugType,accomplishDate,dueDate,bugDelay,iteration) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", trackTitle, trackContent, executor, creator, executor, createTime, updateTime, bugLevel, bugStatus, project, bugType, accomplishDate, dueDate, bugDelay, iteration)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected() //影响行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

func insertUserData(DB *sql.DB, username string, userid string) {
	result, err := DB.Exec("insert INTO user_tb(username, userid) values(?,?)", username, userid)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected() //影响行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

func tokencreate() string {
	//企业id
	//orgId := "5e12a6349cb08f0001fbdbc4"
	//projectId := "5f27ab28ed986600446d197d"

	appId := "6141b2bace049310f8dcd90e"
	secretKey := "jFlz7bLdiVk6uNaUqMaHc1SSTXnF1DKD"

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["_appId"] = appId
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ""
	}
	fmt.Println(tokenString)
	return tokenString
}

func tasksearch() []byte {
	token := tokencreate()

	url := "https://open.teambition.com/api/task/tqlsearch"
	method := "POST"

	payload := strings.NewReader(`{
	"tql": "_projectId=5f27ab28ed986600446d197d AND _creatorId IN (5e37e431fd6e0300011ac478,60652db5b7c7e4065acf8581,5f813e5ddf2fef6db4c97a2e,5fc5dcf4d09b995e3aeaaddc,5ff53e219e3851ca2c1f6454)",
	"pageSize": 3000,
    "pageToken": "",
	"orderBy": "dueDate"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-Tenant-Id", "5e12a6349cb08f0001fbdbc4")
	req.Header.Add("X-Tenant-Type", "organization")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func wywkMember() []byte {
	token := tokencreate()

	url := "https://open.teambition.com/api/org/member/list?orgId=5e12a6349cb08f0001fbdbc4&pageSize=1000"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-Tenant-Id", "5e12a6349cb08f0001fbdbc4")
	req.Header.Add("X-Tenant-Type", "organization")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func projectMember(projectid string) []byte {
	token := tokencreate()

	url := "https://open.teambition.com/api/project/member/list?projectId=" + projectid + "&pageToken&pageSize=1000"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("X-Tenant-Id", "5e12a6349cb08f0001fbdbc4")
	req.Header.Add("X-Tenant-Type", "organization")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func selectMember(DB *sql.DB, userid string) string {
	userbind := new(UserBind)
	row := DB.QueryRow("select username from user_tb where userid=?", userid)
	if err := row.Scan(&userbind.Username); err != nil {
		fmt.Printf("无法匹配员工信息, err:%v", err)
		return "员工已离职"
	}
	return userbind.Username.String
}

func issuesInit(DB *sql.DB) {
	_, err := DB.Exec("TRUNCATE TABLE i")
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
}
func main() {
	tokencreate()
	//body := projectMember("61309876bee0dfac6d17c84e")
	//DB := dbInit()
	//var user UserProject
	//err := json.Unmarshal(body,&user)
	//if err != nil {
	//	fmt.Printf("反序列化错误：%s",err)
	//}
	//fmt.Printf("用户：%+v,userid：%+v",user.Result[0].Name, user.Result[0].UserId)
	////a :=selectMember(DB, user.Result[0].UserId)
	////fmt.Println(a)
	////用户数据存储
	//for index, _ := range user.Result {
	//	username := user.Result[index].Name
	//	userid := user.Result[index].UserId
	//	a :=selectMember(DB, user.Result[index].UserId)
	//	if username != "" && a=="员工已离职"{
	//		fmt.Printf("用户：%+v,userid：%+v",user.Result[index].Name, user.Result[index].UserId)
	//		insertUserData(DB,username,userid)
	//	}
	//}

	//body := tasksearch()
	////fmt.Printf("body 的数据类型是: %T\n",body)
	//var track Track
	//err2 := json.Unmarshal(body, &track)
	////fmt.Printf("%+v",track.Result[0].Customfields)
	//if err2 != nil {
	//	fmt.Printf("反序列化错误：%s", err2)
	//}
	//DB := dbInit()
	//
	//for index, _ := range track.Result {
	//	//缺陷标题
	//	trackTitle := track.Result[index].Content
	//	//创建人
	//	creator := track.Result[index].CreatorId
	//	//缺陷内容
	//	trackContent := track.Result[index].Note
	//	//创建时间
	//	createTime := track.Result[index].Created
	//	//执行人
	//	executor := track.Result[index].ExecutorId
	//	//fmt.Printf("缺陷标题:%s\n，缺陷创建人:%s\n,缺席内容:%s\n,创建时间:%s\n,缺陷等级:%d\n,缺陷状态:%d\n,执行者:%s\n",trackTitle,creator,trackContent,createTime,level,status,rd)
	//	//更新时间
	//	updateTime := track.Result[index].Updated
	//	//缺陷等级
	//	fmt.Printf("测试%+v", track.Result[index])
	//	var bugLevelStr string
	//	var bugType string
	//	var iteration string
	//	cfd := track.Result[index].Customfields
	//	bugLevelStr = "未填写"
	//	bugType = "未填写"
	//	iteration = "未填写"
	//	if len(cfd) > 0 {
	//		for index, _ := range cfd {
	//			if cfd[index].CfId == "5e12a64103bbcc000193b072" {
	//				iteration = cfd[index].Value[0].Title
	//			}
	//			if cfd[index].CfId == "5e12a64103bbcc000193b070" {
	//				bugLevelStr = cfd[index].Value[0].Title
	//			}
	//			if cfd[index].CfId == "5e12a64103bbcc000193b071" {
	//				bugType = cfd[index].Value[0].Title
	//			}
	//			//else {
	//			//	bugLevelStr = "未填写"
	//			//	bugType = "未填写"
	//			//	iteration = "未填写"
	//			//}
	//		}
	//
	//		var bugLevel int
	//		if bugLevelStr == "致命" {
	//			bugLevel = 0
	//		} else if bugLevelStr == "严重" {
	//			bugLevel = 1
	//		} else if bugLevelStr == "一般" {
	//			bugLevel = 2
	//		} else if bugLevelStr == "轻微" {
	//			bugLevel = 3
	//		}
	//		//缺陷状态
	//		bugStatus := track.Result[index].IsDone
	//		//缺陷类型
	//		//完成时间
	//		accomplishDate := track.Result[index].AccomplishDate
	//		//截止时间
	//		dueDate := track.Result[index].DueDate
	//		//是否延期
	//		var bugDelay int
	//		if accomplishDate.Unix() < dueDate.Unix() {
	//			bugDelay = 0
	//		} else {
	//			bugDelay = 1
	//		}
	//		//
	//		project := "核心系统2.0"
	//		creator = selectMember(DB, creator)
	//		executor = selectMember(DB, executor)
	//
	//		insertIssueData(DB, trackTitle, creator, trackContent, createTime, updateTime, index, executor, bugLevel, bugStatus, project, bugType, accomplishDate, dueDate, bugDelay, iteration)
	//	}
}
