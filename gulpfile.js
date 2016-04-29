var gulp = require('gulp');
var less = require('gulp-less');

var files = {
    less: './web/public/less/**/*.less' 
};

gulp.task('less', function() {
    return gulp.src(files.less)
        .pipe(less())
        .pipe(gulp.dest('./web/assets/css'));
})

gulp.task('watch', ['less'], function() {
   gulp.watch(files.less, ['less']); 
});

gulp.task('default', ['less'], function() {
});