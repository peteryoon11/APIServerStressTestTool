# 서비스를 진행 하다 보면 
* 스트레스 테스트가 필요한 경우 
    * 같은 데이터/ 다른 데이터에 대해서  request , respond 를 기록하고 누락되는 데이터가 있는지 확인한다.
* 누락된 데이터를 넣어야 하는 경우 
    * ./MassData 에 있는 파일을 readFile 로 읽어서 하나씩 날리는것을 구현하자.
* 데이터를 전송 했는데 그것에 대한 request , respond 를 기록하기 위한 부분들이 필요하다.  
    * appenFile 을 이용해서 시간과 request, respond를 기록하자.
