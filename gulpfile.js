var gulp = require('gulp'),
	beautify = require('gulp-beautify'),
	jshint = require('gulp-jshint');
 
gulp.task('beautify', function() {
  return gulp.src('./_static/js/*.js')
    .pipe(beautify())
    .pipe(gulp.dest('./_static/js/'))
});

gulp.task('lint', function() {
  return gulp.src('./_static/js/*.js')
    .pipe(jshint());
});

gulp.task('default', ['beautify', 'lint']);