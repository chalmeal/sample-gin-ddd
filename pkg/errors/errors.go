package errors

import "errors"

var (
	// StatusOK(200)
	LOGIN_SUCCESS             = errors.New("ログインしました。")
	TEMPORARY_ACCOUNT_SUCCESS = errors.New("メールアドレスにアカウント登録のURLを送信しました。")
	REGISTER_ACCOUNT_SUCCESS  = errors.New("アカウントを登録しました。")
	EDIT_ACCOUNT_SUCCESS      = errors.New("アカウントを更新しました。")
	DELETE_ACCOUNT_SUCCESS    = errors.New("アカウントを削除しました。")
	REGISTER_TODO_SUCCESS     = errors.New("TODOを登録しました。")
	EDIT_TODO_SUCCESS         = errors.New("TODOを更新しました。")
	DELETE_TODO_SUCCESS       = errors.New("TODOを削除しました。")
	// BadRequest(400)
	LOGIN_FAILURE                      = errors.New("ID又はパスワードを確認してください。")
	TEMPORARY_REGISTER_ACCOUNT_ALREADY = errors.New("既に登録済みのメールアドレスです。")
	TEMPORARY_REGISTER_ACCOUNT_FAILURE = errors.New("メールアドレスを確認してください。")
	REGISTER_ACCOUNT_FAILURE           = errors.New("アカウント登録の入力情報を確認してください。")
	EDIT_ACCOUNT_FAILURE               = errors.New("アカウント編集の入力情報を確認してください。")
	DELETE_ACCOUNT_ALREADY             = errors.New("既に削除済みのアカウントです。")
	REGISTER_TODO_FAILURE              = errors.New("TODO登録の入力情報を確認してください。")
	EDIT_TODO_FAILURE                  = errors.New("TODO編集の入力情報を確認してください。")
	// UnAuthorization(401)
	LOGIN_UN_AUTHORIZATION      = errors.New("ID又はパスワードを確認してください。")
	ACCESS_TOKEN_VERIFY_FAILURE = errors.New("アクセストークンの検証に失敗しました。")
	NOT_ACCESS_AUTHORIZE        = errors.New("アクセスが許可されていません。")
	// NotFound(404)
	GET_ACCOUNT_NOT_FOUND = errors.New("アカウントが見つかりませんでした。")
	GET_TODO_NOT_FOUND    = errors.New("指定したTODOが見つかりませんでした。")
	// InternalServerError(500)
	INTERNAL_SERVER_ERROR = errors.New("サーバー側で問題が発生しました。")
)

var (
	BadRequest = []error{
		LOGIN_FAILURE,
		TEMPORARY_REGISTER_ACCOUNT_ALREADY,
		TEMPORARY_REGISTER_ACCOUNT_FAILURE,
		REGISTER_ACCOUNT_FAILURE,
		EDIT_ACCOUNT_FAILURE,
		DELETE_ACCOUNT_ALREADY,
		REGISTER_TODO_FAILURE,
		EDIT_TODO_FAILURE,
	}
	UnAuthorized = []error{
		LOGIN_UN_AUTHORIZATION,
		ACCESS_TOKEN_VERIFY_FAILURE,
		NOT_ACCESS_AUTHORIZE,
	}
	NotFound = []error{
		GET_ACCOUNT_NOT_FOUND,
		GET_TODO_NOT_FOUND,
	}
)
