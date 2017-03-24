'use strict';
var gulp = require('gulp'),
	beautify = require('gulp-beautify'),
	jshint = require('gulp-jshint'),
	cssbeautify = require('gulp-cssbeautify');
 
gulp.task('beautify', function() {
  return gulp.src('./js/*.js')
    .pipe(beautify())
    .pipe(gulp.dest('./js/'))
});

gulp.task('css', function() {
    return gulp.src('./css/*.css')
        .pipe(cssbeautify())
        .pipe(gulp.dest('./css/'));;
});

gulp.task('lint', function() {
  return gulp.src('./js/*.js')
    .pipe(jshint());
});

gulp.task('default', ['beautify', 'css', 'lint']);