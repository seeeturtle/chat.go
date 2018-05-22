# Chat.go

멀티 플랫폼을 지원하 **(려)** 는 범용적 챗봇 프레임 워크입니다.

## Structure

chat.go 프로젝트의 구조를 설명합니다.

### Scenario

chat.go에서 챗봇의 workflow를 정의할 때 사용되는 단위입니다.

정의는 다음과 같습니다:
```go
type Scenario interface {
    Next(Object) (Scenario, Object)
}
```
단순히 Object를 input으로 받고 다음으로 실행할 Scenario와 Object를 주기만 하면 됩니다.

이런 단순한 특성 덕분에 자유롭게 자신만의 Scenario를 만들고 재사용할 수 있습니다.

chat.go에서 자주 쓰일만한 Scenario들을 모아서 제공할 예정입니다.

지금 구현된 Scenario들은 다음과 같습니다:
* CondScenario:
condition과 behavior를 쌍으로 주어 만약 condition이 참일때, 그에 해당하는
behavior를 실행하게 됩니다.

### Object

chat.go에서 Scenario들간의 정보를 주고 받기 위한 단위입니다.

정의는 다음과 같습니다:
```go
type Object interface {
	MarshalJSON() ([]byte, error)
}
```
Object는 Response를 할때 JSON으로 Marshal하기 위한 MarshalJSON 메서드를 필요로 합니다.

지금 카카오톡의 플러스 친구를 지원하기 위해서 구현된 Object들은 다음과 같습니다:
* Text: string의 alias입니다.
* Keyboard: 네, 문서에 있는 그 Keyboard입니다.
* Message: 메세지를 response할때 사용됩니다.
* Photo: 네, 이 또한 문서에 있는 그 Photo입니다.

### Server

인기있는 웹프레임워크인 [Echo](https://github.com/labstack/echo)에 기반해 있습니다

하지만 직접적인 의존성은 없기 때문에 원한다면 다른 프레임워크로도 자유롭게 바꿔서
사용해도 문제가 없습니다.

현재 chat.go에서 지원하는 기능은 다음과 같습니다:
* 이벤트에 대응해서 Scenario를 설정하여 자동으로 라우팅됨(이 부분은 확실하게 정해진 바가 아닙니다).
* `*echo.Echo`를 반환하여 사용자가 마음대로 커스터마이징할 수 있도록 함.

추후에 지원 **(할)** 기능은 다음과 같습니다:
* 모든 플랫폼에서 대해서 똑같은 인터페이스를 가짐.
* 특정 event을 대해서 Scenario를 설정하면 모든 플랫폼에 자동으로 라우팅이 됨
* Object를 반환하기만 해도 자동으로 플랫폼에 맞게 변형시켜 response함.
