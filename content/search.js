function SearchCtrl($scope, $http) {
  $scope.searchResults = [];

  $scope.search = function() {
    var query = $scope.searchText; 
    $http.get("/api/search/"+query).success(function(data, status, headers, config){
      $scope.searchResults = data;
      console.log(data); 
    });
  };
}
