package application

import "github.com/sixfwa/fiber-api/domain/repository"

type UserUsercase struct {
	user repository.UserInterface
}

// Controller
func (uc *userController) GetUser(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("userid")
	if err != nil {
		return FiberFailedParamError("userid", c.Params("userid"))
	}

	// Call business layer with business logic request
	user, err := uc.userUsecase.GetUserByID(c.Context(), userId)
	if err != nil {
		return err
	}

	response := adapters.ToGetUserResponse(user)
	return c.Status(fiber.StatusOK).JSON(response)
}

// Application
func (a *userUsecase) GetUserByID(ctx context.Context, userID int) (*entities.User, error) {
	user, err := a.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, apperr.New("user not found", apperr.StatusNotFound, err)
	}
	return user, nil
}

// Database
func (r *userReadRepositoryPostgres) GetByID(ctx context.Context, id int) (*entities.User, error) {
	var userRow UserRow
	err := r.DB.WithContext(ctx).First(&userRow, id).Error
	if err != nil {
		log.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find user with id %d", id)
		}
		return nil, err
	}

	var postCount int64
	r.DB.WithContext(ctx).Model(&PostRow{}).Where("user_id = ?", id).Count(&postCount)

	var commentCount int64
	r.DB.WithContext(ctx).Model(&CommentRow{}).Where("user_id = ?", id).Count(&commentCount)

	userRow.PostCount = int(postCount)
	userRow.CommentCount = int(commentCount)

	return userRow.ToDomain(), nil
}


func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}