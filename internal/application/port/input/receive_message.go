package input

import (
	"context"
)

// ReceiveMessageUseCase 메시지 수신 유스케이스.
// inputID는 config에 정의된 input의 ID (예: "beszel")로 전달된다.
// contentType은 HTTP Content-Type 헤더 값이다 (파서 선택에 사용).
// 성공 시 생성된 message ID를 반환한다.
type ReceiveMessageUseCase interface {
	Receive(ctx context.Context, inputID string, contentType string, body []byte) (string, error)
}
