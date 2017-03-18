var gulp = require('gulp'),
	beautify = require('gulp-beautify'),
	jshint = require('gulp-jshint');
 
gulp.task('beautify', function() {
  return gulp.src('./static/js/*.js')
    .pipe(beautify())
    .pipe(gulp.dest('./static/js/'))
});

gulp.task('lint', function() {
  return gulp.src('./static/js/*.js')
    .pipe(jshint());
});

gulp.task('default', ['beautify', 'lint']);