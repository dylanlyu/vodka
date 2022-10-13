package member_phone_test

import (
	model "app.inherited.magic/internal/entity/db/members_phone"
	"app.inherited.magic/internal/entity/member_phone"
	"app.inherited.magic/internal/interactor/util"
	dbConfig "app.inherited.magic/internal/interactor/util/connect"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"time"
)

var _ = Describe("Member phone entity", func() {
	var (
		db     *gorm.DB
		mockDB *sql.DB
		mock   sqlmock.Sqlmock
		err    error
		input  *model.Base
		query  string
	)

	BeforeEach(func() {
		mockDB, mock, err = sqlmock.New()
		Expect(err).Should(Succeed())
		config := dbConfig.PostgresConfig{}
		config.DSN = util.PointerString("sqlmock_db_0")
		config.DriverName = util.PointerString("postgres")
		config.Conn = mockDB
		config.PreferSimpleProtocol = util.PointerBool(true)
		config.NowFunc = func() time.Time { return time.Now().UTC() }
		config.Logger = logger.Default.LogMode(logger.Info)
		db, err = config.Connect()
		Expect(err).Should(Succeed())
		input = &model.Base{}
	})

	Describe("Member phone create", func() {
		JustBeforeEach(func() {
			input.ID = util.PointerString("106d0c5c-053c-44c0-9bf7-7524086c4909")
			input.PhoneCode = util.PointerString("886")
			input.PhoneNumber = util.PointerString("910213888")
			input.CreatedAt = util.PointerTime(util.NowToUTC())
			input.UpdatedAt = util.PointerTime(util.NowToUTC())
			query = `INSERT INTO "members_phone" ("phone_code","phone_number","created_at","updated_at","deleted_at","id") 
							VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`
		})

		When("New member create", func() {
			It("Success", func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(input.PhoneCode, input.PhoneNumber,
					input.CreatedAt, input.UpdatedAt, input.DeletedAt, input.ID).
					WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(input.ID))
				mock.ExpectCommit()
				storage := member_phone.Init(db)
				Expect(storage.Create(input)).Should(Succeed())
			})

			It("Error", func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(input.PhoneCode, input.PhoneNumber,
					input.CreatedAt, input.UpdatedAt, input.DeletedAt, input.ID).
					WillReturnError(errors.
						New(`[23505] ERROR: duplicate key value violates unique constraint "members_phone_pkey"`))
				mock.ExpectRollback()
				storage := member_phone.Init(db)
				Expect(storage.Create(input)).Should(
					MatchError(`[23505] ERROR: duplicate key value violates unique constraint "members_phone_pkey"`))
			})
		})
	})
})
