package bdfConv

import "C"
import "unsafe"

/*
#include <string.h>
#include <stdlib.h>
#include <time.h>
#include "bdf_font.h"
#include "fd.h"


int get_str_arg(char ***argv, int c, char **result)
{
  if ( (**argv)[0] == '-' )
  {
    if ( (**argv)[1] == c )
    {
      if ( (**argv)[2] == '\0' )
      {
	(*argv)++;
	*result = **argv;
      }
      else
      {
	*result = (**argv)+2;
      }
      (*argv)++;
      return 1;
    }
  }
  return 0;
}


int get_num_arg(char ***argv, int c, unsigned long *result)
{
  if ( (**argv)[0] == '-' )
  {
    if ( (**argv)[1] == c )
    {

      if ( (**argv)[2] == '\0' )
      {
	if ( *((*argv)+1) != NULL )
	{
	  (*argv)++;
	  *result = strtoul((**argv), NULL, 10);
	}
      }
      else
      {
	*result = strtoul((**argv)+2, NULL, 10);
      }
      (*argv)++;
      return 1;
    }
  }
  return 0;
}

int get_num_strarg(char ***argv, const char *s, unsigned long *result)
{
  if ( (**argv)[0] == '-' )
  {
    //printf("get_num_strarg %s: match %s\n", **argv, s);
    if ( strcmp( (**argv)+1, s ) == 0 )
    {
      //printf("get_num_strarg %s: match %s found\n", **argv, s);
      if ( *((*argv)+1) != NULL )
      {
	(*argv)++;
	*result = strtoul((**argv), NULL, 10);
      }
      (*argv)++;
      return 1;
    }
  }
  return 0;
}

int is_arg(char ***argv, int c)
{
  if ( (**argv)[0] == '-' )
  {
    if ( (**argv)[1] == c )
    {
      (*argv)++;
      return 1;
    }
  }
  return 0;
}

unsigned long left_margin = 1;
unsigned long build_bbx_mode = 0;
unsigned long font_format = 0;
unsigned long min_distance_in_per_cent_of_char_width = 25;
unsigned long cmdline_glyphs_per_line = 16;
unsigned long xoffset = 0;
unsigned long yoffset = 0;
unsigned long tile_h_size = 1;
unsigned long tile_v_size = 1;
int font_picture_extra_info = 0;
int font_picture_test_string = 0;
int runtime_test = 0;
char *c_filename = NULL;
char *k_filename = NULL;
char *target_fontname = "bdf_font";

unsigned tga_get_line_height(bf_t *bf_desc_font, bf_t *bf)
{
  unsigned h;
  tga_set_font(bf_desc_font->target_data);
  h = tga_get_char_height();
  tga_set_font(bf->target_data);
  if ( h < tga_get_char_height() )
    return tga_get_char_height();
  return h;
}

unsigned tga_draw_font_line(unsigned y, long enc_start, bf_t *bf_desc_font, bf_t *bf, long glyphs_per_line)
{
  long i;
  unsigned x;
  int is_empty;
  char pre[32];

  is_empty = 1;
  for( i = 0; i< 16 && is_empty != 0; i++ )
  {
    if ( tga_get_glyph_data(i+enc_start) != NULL )
	is_empty = 0;
  }

  if ( is_empty != 0 )
    return 0;

  sprintf(pre, "%5ld/0x%04lx", enc_start, enc_start);

  x = left_margin;
  if ( bf_desc_font != NULL )
  {
    if ( bf_desc_font->target_data != NULL )
    {
      tga_set_font(bf_desc_font->target_data);
      x += tga_draw_string(x, y, pre, 0, 0);
    }
  }
  x += 4;

  tga_set_font(bf->target_data);
  for( i = 0; i< glyphs_per_line; i++ )
  {
    tga_draw_glyph(x + (tga_get_char_width()+2)*i,y,enc_start+i, font_picture_extra_info);
  }

  return left_margin + x + (tga_get_char_width()+2)*glyphs_per_line;
}

unsigned tga_draw_font_info(unsigned y, const char *fontname, bf_t *bf_desc_font, bf_t *bf)
{
  unsigned x;
  int cap_a, cap_a_height;
  static char s[256];

  cap_a_height = 0;
  cap_a = bf_GetIndexByEncoding(bf, 'A');
  if ( cap_a >= 0 )
  {
    cap_a_height = bf->glyph_list[cap_a]->bbx.h+bf->glyph_list[cap_a]->bbx.y;
  }

  if ( bf_desc_font != NULL )
  {
    if ( bf_desc_font->target_data != NULL )
    {

      tga_set_font(bf_desc_font->target_data);

      y +=  tga_get_char_height()+1;
      x = left_margin;
      x += tga_draw_string(x, y, fontname, 0, 0);

      y +=  tga_get_char_height()+1;
      sprintf(s, "Max width %u, max height %u", tga_get_char_width(), tga_get_char_height());
      x = left_margin;
      x += tga_draw_string(x, y, s, 0, 0);

      y +=  tga_get_char_height()+1;
      sprintf(s, "'A' height %d, font size %d ", cap_a_height, bf->target_cnt);
      x = left_margin;
      x += tga_draw_string(x, y, s, 0, 0);
      return (tga_get_char_height()+1)*3;
    }
  }
  return 0;
}


unsigned tga_draw_font(unsigned y, const char *fontname, bf_t *bf_desc_font, bf_t *bf, long glyphs_per_line)
{
  long i;
  unsigned x, xmax;
  xmax = 0;

  bf_Log(bf, "Draw TGA, line height %d", tga_get_line_height(bf_desc_font, bf));

  y += tga_draw_font_info( y, fontname, bf_desc_font, bf);

  y +=   tga_get_line_height(bf_desc_font, bf)+1;



  for( i = 0; i <= 0x0ffff; i+=glyphs_per_line )
  {
    x = tga_draw_font_line(y, i, bf_desc_font, bf, glyphs_per_line);
    if ( x > 0 )
    {
      if ( xmax < x )
	xmax = x;
      y +=  tga_get_line_height(bf_desc_font, bf)+1;
    }
  }

  bf_Log(bf, "Draw TGA, xmax %d", xmax);

  tga_set_font(bf->target_data);

  if ( font_picture_test_string != 0 )
  {
    tga_draw_string(left_margin, y, "Woven silk pyjamas exchanged for blue quartz.", 0, xmax);
    y +=  tga_get_line_height(bf_desc_font, bf)+1;
  }
  return y;
}


int conv(char* _bdf_filename, char* _map_str, char* _target_fontname, char* _c_filename) {
	bf_t *bf;
	char *bdf_filename = _bdf_filename;
	int is_verbose = 0;
	char *map_str = "*";
	char *map_filename ="";
	char *desc_font_str = "";
	unsigned y;

	target_fontname = _target_fontname;
	c_filename = _c_filename;
	map_str = _map_str;

	bf = bf_OpenFromFile(bdf_filename, is_verbose, build_bbx_mode, map_str, map_filename, font_format, xoffset, yoffset, tile_h_size, tile_v_size);

	if ( bf == NULL ){
		return -1;
	}
	bf_WriteU8G2CByFilename(bf, c_filename, target_fontname, "  ");
	bf_Close(bf);
	return 0;
}
*/
import "C"

// go wrapper for the bdfconv tools from u8g2 source code

func CreateTargetCFont(sourceBdfFilePath, mapStr, exportFontName, exportCFilePath string) int {
	if mapStr == "" {
		mapStr = "*"
	}
	cSourcebdffilepath := C.CString(sourceBdfFilePath)
	defer C.free(unsafe.Pointer(cSourcebdffilepath))

	cMapStr := C.CString(mapStr)
	defer C.free(unsafe.Pointer(cMapStr))

	cExportfontname := C.CString(exportFontName)
	defer C.free(unsafe.Pointer(cExportfontname))

	cExportcfilepath := C.CString(exportCFilePath)
	defer C.free(unsafe.Pointer(cExportcfilepath))
	cResult := C.conv(cSourcebdffilepath, cMapStr, cExportfontname, cExportcfilepath)
	return int(cResult)
}
